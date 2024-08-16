package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/waydxd/paopao-ce/internal/model/web"
)

func init() {
	Entry[Community]()
}

type Community struct {
	Chain `mir:"-"`
	Group `mir:"v1"`

	CreateCommunity func(Post, web.CreateCommunityReq) web.CreateCommunityResp `mir:"/community/create"`

	GetCommunity func(Get, web.GetCommunityReq) web.GetCommunityResp `mir:"/community/get"`

	ListCommunities func(Get, web.ListCommunitiesReq) web.ListCommunitiesResp `mir:"/community/list"`

	JoinCommunity func(Post, web.JoinLeaveCommunityReq) `mir:"/community/join"`

	LeaveCommunity func(Post, web.JoinLeaveCommunityReq) `mir:"/community/leave"`

	ListCommunityMembers func(Get, web.ListCommunityMembersReq) web.ListCommunityMembersResp `mir:"/community/members"`

	GetCommunityPost func(Get, web.GetCommunityPostReq) web.GetCommunityPostResp `mir:"/community/posts"`
}
