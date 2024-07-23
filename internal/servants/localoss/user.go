// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package localoss

import (
	api "github.com/waydxd/paopao-ce/auto/api/s/v1"
	"github.com/waydxd/paopao-ce/internal/servants/base"
)

var (
	_ api.User = (*userSrv)(nil)
)

type userSrv struct {
	*base.BaseServant
	api.UnimplementedUserServant
}

func newUserSrv() api.User {
	return &userSrv{
		BaseServant: base.NewBaseServant(),
	}
}
