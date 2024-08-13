package core

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
)

// interface CommunityModelService
type CommunityModelService interface {
	CreateCommunity(community *ms.Community) error
	RemoveCommunity(communityID int64) error
	AddMember(userID, communityID int64) error
	RemoveMember(userID, communityID int64) error
	JoinCommunity(userID, communityID int64) error
	LeaveCommunity(userID, communityID int64) error
}
type CommunityService interface {
	CreateCommunity(name, description string, creatorID int64) (*ms.Community, error)
	GetCommunity(communityID int64) (*ms.Community, error)
	ListCommunities(offset, limit int) ([]*ms.Community, error)
	ListCommunityMembers(communityID int64, offset, limit int) ([]*ms.User, error)
}
