// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/waydxd/paopao-ce/internal/model/web"
)

type _binding_ interface {
	Bind(*gin.Context) mir.Error
}

type _render_ interface {
	Render(*gin.Context)
}

type _default_ interface {
	Bind(*gin.Context, any) mir.Error
	Render(*gin.Context, any, mir.Error)
}

type Admin interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	SiteInfo(*web.SiteInfoReq) (*web.SiteInfoResp, mir.Error)
	ChangeUserStatus(*web.ChangeUserStatusReq) mir.Error

	mustEmbedUnimplementedAdminServant()
}

// RegisterAdminServant register Admin servant to gin
func RegisterAdminServant(e *gin.Engine, s Admin) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/admin/site/status", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.SiteInfoReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.SiteInfo(req)
		s.Render(c, resp, err)
	})
	router.Handle("POST", "/admin/user/status", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.ChangeUserStatusReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.ChangeUserStatus(req))
	})
}

// UnimplementedAdminServant can be embedded to have forward compatible implementations.
type UnimplementedAdminServant struct{}

func (UnimplementedAdminServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedAdminServant) SiteInfo(req *web.SiteInfoReq) (*web.SiteInfoResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) ChangeUserStatus(req *web.ChangeUserStatusReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAdminServant) mustEmbedUnimplementedAdminServant() {}
