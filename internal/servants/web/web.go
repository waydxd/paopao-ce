// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"sync"

	"github.com/alimy/tryst/cfg"
	"github.com/gin-gonic/gin"
	api "github.com/waydxd/paopao-ce/auto/api/v1"
	"github.com/waydxd/paopao-ce/internal/conf"
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/dao"
	"github.com/waydxd/paopao-ce/internal/dao/cache"
	"github.com/waydxd/paopao-ce/internal/servants/base"
)

var (
	_enablePhoneVerify    bool
	_disallowUserRegister bool
	_ds                   core.DataService
	_ac                   core.AppCache
	_wc                   core.WebCache
	_oss                  core.ObjectStorageService
	_onceInitial          sync.Once
)

// RouteWeb register web route
func RouteWeb(e *gin.Engine) {
	lazyInitial()
	ds := base.NewDaoServant()
	// aways register servants
	api.RegisterAdminServant(e, newAdminSrv(ds, _wc))
	api.RegisterCoreServant(e, newCoreSrv(ds, _oss, _wc))
	api.RegisterRelaxServant(e, newRelaxSrv(ds, _wc), newRelaxChain())
	api.RegisterLooseServant(e, newLooseSrv(ds, _ac))
	api.RegisterPrivServant(e, newPrivSrv(ds, _oss), newPrivChain())
	api.RegisterPubServant(e, newPubSrv(ds))
	api.RegisterTrendsServant(e, newTrendsSrv(ds))
	api.RegisterFollowshipServant(e, newFollowshipSrv(ds))
	api.RegisterFriendshipServant(e, newFriendshipSrv(ds))
	api.RegisterSiteServant(e, newSiteSrv())
	// regster servants if needed by configure
	cfg.Be("Alipay", func() {
		client := conf.MustAlipayClient()
		api.RegisterAlipayPubServant(e, newAlipayPubSrv(ds))
		api.RegisterAlipayPrivServant(e, newAlipayPrivSrv(ds, client))
	})
	// shedule jobs if need
	scheduleJobs()
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		_enablePhoneVerify = cfg.If("Sms")
		_disallowUserRegister = cfg.If("Web:DisallowUserRegister")
		_maxWhisperNumDaily = conf.AppSetting.MaxWhisperDaily
		_maxCaptchaTimes = conf.AppSetting.MaxCaptchaTimes
		_oss = dao.ObjectStorageService()
		_ds = dao.DataService()
		_ac = cache.NewAppCache()
		_wc = cache.NewWebCache()
	})
}
