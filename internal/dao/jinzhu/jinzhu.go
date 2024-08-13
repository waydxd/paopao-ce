// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// package jinzhu Core service implement base gorm+mysql/postgresql/sqlite3.
// Jinzhu is the primary developer of gorm so use his name as
// package name as a saluter.

package jinzhu

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/waydxd/paopao-ce/internal/conf"
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/dao/cache"
	"github.com/waydxd/paopao-ce/internal/dao/security"
)

var (
	_ core.DataService = (*dataSrv)(nil)
	_ core.VersionInfo = (*dataSrv)(nil)

	_ core.WebDataServantA = (*webDataSrvA)(nil)
	_ core.VersionInfo     = (*webDataSrvA)(nil)

	_onceInitial sync.Once
)

type dataSrv struct {
	core.WalletService
	core.MessageService
	core.TopicService
	core.TweetService
	core.TweetManageService
	core.TweetHelpService
	core.TweetMetricServantA
	core.CommentService
	core.CommentManageService
	core.CommentMetricServantA
	core.TrendsManageServantA
	core.UserManageService
	core.UserMetricServantA
	core.ContactManageService
	core.FollowingManageService
	core.UserRelationService
	core.SecurityService
	core.AttachmentCheckService
	core.CommunityService
	core.CommunityManageService
}

type webDataSrvA struct {
	core.TopicServantA
	core.TweetServantA
	core.TweetManageServantA
	core.TweetHelpServantA
}

func NewDataService() (core.DataService, core.VersionInfo) {
	lazyInitial()
	db := conf.MustGormDB()
	pvs := security.NewPhoneVerifyService()
	tms := newTweetMetricServentA(db)
	ums := newUserMetricServentA(db)
	cms := newCommentMetricServentA(db)
	cis := cache.NewEventCacheIndexSrv(tms)
	ds := &dataSrv{
		TweetMetricServantA:    tms,
		CommentMetricServantA:  cms,
		UserMetricServantA:     ums,
		WalletService:          newWalletService(db),
		MessageService:         newMessageService(db),
		TopicService:           newTopicService(db),
		TweetService:           newTweetService(db),
		TweetManageService:     newTweetManageService(db, cis),
		TweetHelpService:       newTweetHelpService(db),
		CommentService:         newCommentService(db),
		CommentManageService:   newCommentManageService(db),
		TrendsManageServantA:   newTrendsManageServentA(db),
		UserManageService:      newUserManageService(db, ums),
		ContactManageService:   newContactManageService(db),
		FollowingManageService: newFollowingManageService(db),
		UserRelationService:    newUserRelationService(db),
		SecurityService:        newSecurityService(db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
		CommunityService:       newCommunityService(db),
		CommunityManageService: newCommunityManageService(db),
	}
	return cache.NewCacheDataService(ds), ds
}

func NewWebDataServantA() (core.WebDataServantA, core.VersionInfo) {
	lazyInitial()
	db := conf.MustGormDB()
	ds := &webDataSrvA{
		TopicServantA:       newTopicServantA(db),
		TweetServantA:       newTweetServantA(db),
		TweetManageServantA: newTweetManageServantA(db),
		TweetHelpServantA:   newTweetHelpServantA(db),
	}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return newAuthorizationManageService(conf.MustGormDB())
}

func (s *dataSrv) Name() string {
	return "Gorm"
}

func (s *dataSrv) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}

func (s *webDataSrvA) Name() string {
	return "Gorm"
}

func (s *webDataSrvA) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		initTableName()
	})
}
