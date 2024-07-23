// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v4"
	api "github.com/waydxd/paopao-ce/auto/api/v1"
	"github.com/waydxd/paopao-ce/internal/conf"
	"github.com/waydxd/paopao-ce/internal/model/web"
	"github.com/waydxd/paopao-ce/internal/servants/base"
	"github.com/waydxd/paopao-ce/pkg/version"
)

var (
	_ api.Site = (*siteSrv)(nil)
)

type siteSrv struct {
	api.UnimplementedSiteServant
	*base.BaseServant
}

func (*siteSrv) Profile() (*web.SiteProfileResp, mir.Error) {
	return conf.WebProfileSetting, nil
}

func (*siteSrv) Version() (*web.VersionResp, mir.Error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

func newSiteSrv() api.Site {
	return &siteSrv{
		BaseServant: base.NewBaseServant(),
	}
}
