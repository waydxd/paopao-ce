package web

import (
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/waydxd/paopao-ce/auto/api/v1"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/model/web"
	"github.com/waydxd/paopao-ce/internal/servants/base"
	"github.com/waydxd/paopao-ce/internal/servants/chain"
)

var _ api.Community = (*communitySrv)(nil)

type communitySrv struct {
	api.UnimplementedCommunityServant
	*base.DaoServant
}

func newCommunitySrv(s *base.DaoServant) *communitySrv {
	return &communitySrv{
		DaoServant: s,
	}
}

func (c *communitySrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (c *communitySrv) ListCommunityMembers(req *web.ListCommunityMembersReq) (*web.ListCommunityMembersResp, mir.Error) {
	members, err := c.Ds.ListCommunityMembers(uint(req.CommunityId), 0, 20) // Using default offset and limit
	if err != nil {
		return nil, web.ErrListCommunity
	}
	return &web.ListCommunityMembersResp{Members: members}, nil
}

func (c *communitySrv) LeaveCommunity(req *web.JoinLeaveCommunityReq) mir.Error {
	err := c.Ds.LeaveCommunity(0, uint(req.CommunityId)) // Assuming userID is 0 for now
	if err != nil {
		return web.ErrLeaveCommunity
	}
	return nil
}

func (c *communitySrv) JoinCommunity(req *web.JoinLeaveCommunityReq) mir.Error {
	err := c.Ds.JoinCommunity(0, uint(req.CommunityId)) // Assuming userID is 0 for now
	if err != nil {
		return web.ErrJoinCommunity
	}
	return nil
}

func (c *communitySrv) ListCommunities(req *web.ListCommunitiesReq) (*web.ListCommunitiesResp, mir.Error) {
	communities, err := c.Ds.ListCommunities(req.Offset, req.Limit) // Using default offset and limit
	if err != nil {
		return nil, web.ErrListCommunity
	}
	return &web.ListCommunitiesResp{Communities: communities}, nil
}

func (c *communitySrv) GetCommunity(req *web.GetCommunityReq) (*web.GetCommunityResp, mir.Error) {
	community, err := c.Ds.GetCommunity(uint(req.CommunityId))
	if err != nil {
		return nil, web.ErrGetCommunity
	}
	return &web.GetCommunityResp{Community: community}, nil
}

func (c *communitySrv) CreateCommunity(req *web.CreateCommunityReq) (*web.CreateCommunityResp, mir.Error) {
	community := &ms.Community{
		Name:        req.Name,
		Description: req.Description,
		AvatarURL:   req.AvatarURL,
		BannerURL:   req.BannerURL,
	}

	err := c.Ds.CreateCommunity(community)
	if err != nil {
		return nil, web.ErrCreateCommunity
	}

	// At this point, 'community' should have an ID assigned by the database

	return &web.CreateCommunityResp{Community: community}, nil
}
