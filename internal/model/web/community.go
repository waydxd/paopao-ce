package web

import (
	"github.com/waydxd/paopao-ce/internal/dao/jinzhu/dbr"
)

type JoinLeaveCommunityReq struct {
	CommunityId int64 `json:"community_id" binding:"required"`
}
type CreateCommunityReq struct {
	dbr.Community
}
type CreateCommunityResp struct {
}
type ListCommunitiesReq struct {
}
type ListCommunitiesResp struct {
}
type ListCommunityMembersReq struct {
	CommunityId int64 `json:"community_id" binding:"required"`
}
type ListCommunityMembersResp struct {
}
type GetCommunityReq struct {
}
type GetCommunityResp struct {
}
