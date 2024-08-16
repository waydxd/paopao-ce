package core

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
)

// interfaces

type CommunityManageService interface {
	CreateCommunity(community *ms.Community) error
	RemoveCommunity(communityID uint) error
	AddMember(userID, communityID uint) error
	RemoveMember(userID, communityID uint) error
	JoinCommunity(userID, communityID uint) error
	LeaveCommunity(userID, communityID uint) error
}
type CommunityService interface {
	GetCommunity(communityID uint) (*ms.Community, error)
	ListCommunities(offset, limit int) ([]*ms.Community, error)
	ListCommunityMembers(communityID uint, offset int, limit int) ([]*ms.User, error)
	GetCommunityPost(communityID uint, offset int, limit int) (*ms.IndexTweetList, error)
}
