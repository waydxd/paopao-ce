// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/waydxd/paopao-ce/internal/conf"
	"github.com/waydxd/paopao-ce/pkg/version"
)

type VersionResp struct {
	BuildInfo *version.BuildInfo `json:"build_info"`
}

type SiteProfileResp = conf.WebProfileConf
