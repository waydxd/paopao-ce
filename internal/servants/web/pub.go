// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/afocus/captcha"
	"github.com/alimy/mir/v4"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	api "github.com/waydxd/paopao-ce/auto/api/v1"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/model/web"
	"github.com/waydxd/paopao-ce/internal/servants/base"
	"github.com/waydxd/paopao-ce/internal/servants/web/assets"
	"github.com/waydxd/paopao-ce/pkg/app"
	"github.com/waydxd/paopao-ce/pkg/utils"
	"github.com/waydxd/paopao-ce/pkg/version"
	"github.com/waydxd/paopao-ce/pkg/xerror"
)

var (
	_ api.Pub = (*pubSrv)(nil)
)

const (
	_MaxLoginErrTimes = 10
	_MaxPhoneCaptcha  = 10
)

type pubSrv struct {
	api.UnimplementedPubServant
	*base.DaoServant
}

func (s *pubSrv) SendCaptcha(req *web.SendCaptchaReq) mir.Error {
	ctx := context.Background()

	// 验证图片验证码
	if captcha, err := s.Redis.GetImgCaptcha(ctx, req.ImgCaptchaID); err != nil || string(captcha) != req.ImgCaptcha {
		logrus.Debugf("get captcha err:%s expect:%s got:%s", err, captcha, req.ImgCaptcha)
		return web.ErrErrorCaptchaPassword
	}
	s.Redis.DelImgCaptcha(ctx, req.ImgCaptchaID)

	// 今日频次限制
	if count, _ := s.Redis.GetCountSmsCaptcha(ctx, req.Phone); count >= _MaxPhoneCaptcha {
		return web.ErrTooManyPhoneCaptchaSend
	}

	if err := s.Ds.SendPhoneCaptcha(req.Phone); err != nil {
		return xerror.ServerError
	}
	// 写入计数缓存
	s.Redis.IncrCountSmsCaptcha(ctx, req.Phone)

	return nil
}

func (s *pubSrv) GetCaptcha() (*web.GetCaptchaResp, mir.Error) {
	cap := captcha.New()
	if err := cap.AddFontFromBytes(assets.ComicBytes); err != nil {
		logrus.Errorf("cap.AddFontFromBytes err:%s", err)
		return nil, xerror.ServerError
	}
	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	if err := png.Encode(emptyBuff, img); err != nil {
		logrus.Errorf("png.Encode err:%s", err)
		return nil, xerror.ServerError
	}
	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())
	// 五分钟有效期
	s.Redis.SetImgCaptcha(context.Background(), key, password)
	return &web.GetCaptchaResp{
		Id:      key,
		Content: "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	}, nil
}

func (s *pubSrv) Register(req *web.RegisterReq) (*web.RegisterResp, mir.Error) {
	if _disallowUserRegister {
		return nil, web.ErrDisallowUserRegister
	}
	// 用户名检查
	if err := s.validUsername(req.Username); err != nil {
		return nil, err
	}
	// 密码检查
	if err := checkPassword(req.Password); err != nil {
		logrus.Errorf("scheckPassword err: %v", err)
		return nil, web.ErrUserRegisterFailed
	}
	password, salt := encryptPasswordAndSalt(req.Password)
	user := &ms.User{
		Nickname: req.Username,
		Username: req.Username,
		Password: password,
		Avatar:   getRandomAvatar(),
		Salt:     salt,
		Status:   ms.UserStatusNormal,
	}
	user, err := s.Ds.CreateUser(user)
	if err != nil {
		logrus.Errorf("Ds.CreateUser err: %s", err)
		return nil, web.ErrUserRegisterFailed
	}
	return &web.RegisterResp{
		UserId:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *pubSrv) Login(req *web.LoginReq) (*web.LoginResp, mir.Error) {
	ctx := context.Background()
	user, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err:%s", err)
		return nil, xerror.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if count, err := s.Redis.GetCountLoginErr(ctx, user.ID); err == nil && count >= _MaxLoginErrTimes {
			return nil, web.ErrTooManyLoginError
		}
		// 对比密码是否正确
		if validPassword(user.Password, req.Password, user.Salt) {
			if user.Status == ms.UserStatusClosed {
				return nil, web.ErrUserHasBeenBanned
			}
			// 清空登录计数
			s.Redis.DelCountLoginErr(ctx, user.ID)
		} else {
			// 登录错误计数
			s.Redis.IncrCountLoginErr(ctx, user.ID)
			return nil, xerror.UnauthorizedAuthFailed
		}
	} else {
		return nil, xerror.UnauthorizedAuthNotExist
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		return nil, xerror.UnauthorizedTokenGenerate
	}
	return &web.LoginResp{
		Token: token,
	}, nil
}

func (s *pubSrv) Version() (*web.VersionResp, mir.Error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

// validUsername 验证用户
func (s *pubSrv) validUsername(username string) mir.Error {
	// Commented because authentication function is partially delegated to the front end
	// 检测用户是否合规
	// if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
	// 	return web.ErrUsernameLengthLimit
	// }
	//
	// if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
	// 	return web.ErrUsernameCharLimit
	// }

	// 重复检查
	user, _ := s.Ds.GetUserByUsername(username)
	if user.Model != nil && user.ID > 0 {
		return web.ErrUsernameHasExisted
	}
	return nil
}

func newPubSrv(s *base.DaoServant) api.Pub {
	return &pubSrv{
		DaoServant: s,
	}
}

func (s *pubSrv) CookieLogin(req *web.CheckCookieReq) (*web.LoginResp, mir.Error) {

	KeyPath := "public.pem" // Path to the asymmetric key

	fmt.Printf(req.Token)
	// Decode the base64url encoded key
	KeyData, err := os.ReadFile(KeyPath)
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
		return nil, xerror.ServerError
	}
	fmt.Printf("Decoded secret key: %x\n", KeyData)
	// Parse the private key
	Key, err := jwt.ParseECPublicKeyFromPEM(KeyData)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}
	// Extract the JWT token from the request
	tokenString := req.Token
	fmt.Printf("Token string: %s\n", tokenString)
	// Parse and verify the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the private key
		return Key, nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return nil, xerror.BadCookie
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user information from claims
		fmt.Printf("JWT: %v\n", claims)
		username := claims["data"].(map[string]interface{})["username"].(string) // Create a LoginReq object
		loginReq := &web.LoginReq{
			Username: username,
			Password: "password", // Password is not required for cookie login
		}
		// Attempt to login
		loginResp, loginErr := s.Login(loginReq)
		if loginErr != nil {
			// If login fails, attempt to register the user
			registerReq := &web.RegisterReq{
				Username: username,
				Password: "password", // Password is not required for cookie login
			}
			registerResp, registerErr := s.Register(registerReq)
			if registerErr != nil {
				return nil, registerErr
			}
			loginReq := &web.LoginReq{
				Username: registerResp.Username,
				Password: "password", // Password is not required for cookie login
			}
			loginResp, loginErr := s.Login(loginReq)
			if loginErr != nil {
				// If registration is successful, return the registration response
				return loginResp, nil
			}
		}
		// If login is successful, return the login response
		return loginResp, nil
	} else {
		return nil, xerror.UnauthorizedAuthFailed
	}
}
