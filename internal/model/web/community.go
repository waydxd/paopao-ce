package web

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
)

type JoinLeaveCommunityReq struct {
	CommunityId int64 `json:"community_id" binding:"required"`
}

type CreateCommunityReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatar_url"`
	BannerURL   string `json:"banner_url"`
}

type CreateCommunityResp struct {
	Community *ms.Community `json:"community"`
}

type ListCommunitiesReq struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit,default=20"`
}

type ListCommunitiesResp struct {
	Communities []*ms.Community `json:"communities"`
	Total       int             `json:"total"`
}

type ListCommunityMembersReq struct {
	CommunityId int64 `form:"community_id" binding:"required"`
	Offset      int   `form:"offset,default=0"`
	Limit       int   `form:"limit,default=20"`
}

type ListCommunityMembersResp struct {
	Members []*ms.User `json:"members"`
}

type GetCommunityReq struct {
	CommunityId int64 `form:"community_id" binding:"required"`
}

type GetCommunityResp struct {
	Community *ms.Community `json:"community"`
}

type GetCommunityPostReq struct {
	CommunityId int64 `form:"community_id" binding:"required"`
	Offset      int   `form:"offset,default=0"`
	Limit       int   `form:"limit,default=20"`
}

type GetCommunityPostResp struct {
	Posts []*ms.Post `json:"posts"`
}
