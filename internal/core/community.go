package core

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
)

type CommunityModelService interface {
	CreateCommunity(community *ms.Community) error
	GetCommunity(communityID int64) (*ms.Community, error)
	ListCommunities(offset, limit int) ([]*ms.Community, error)
	AddMember(userID, communityID int64) error
	RemoveMember(userID, communityID int64) error
	ListMembers(communityID int64, offset, limit int) ([]*ms.User, error)
}
