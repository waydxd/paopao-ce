package jinzhu

import (
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type communityDAO struct {
	db *gorm.DB
}

// ListCommunityMembers implements core.CommunityService.

func NewCommunityService(db *gorm.DB) core.CommunityService {
	return &communityDAO{db: db}
}

func (d *communityDAO) CreateCommunity(community *ms.Community) error {
	return d.db.Create(community).Error
}

func (d *communityDAO) GetCommunity(communityID int64) (*ms.Community, error) {
	var community ms.Community
	if err := d.db.First(&community, communityID).Error; err != nil {
		return nil, err
	}
	return &community, nil
}

func (d *communityDAO) ListCommunities(offset, limit int) ([]*ms.Community, error) {
	var communities []*ms.Community
	if err := d.db.Offset(offset).Limit(limit).Find(&communities).Error; err != nil {
		return nil, err
	}
	return communities, nil
}
func (d *communityDAO) ListCommunityMembers(communityID int64, offset int, limit int) ([]*dbr.User, error) {
	var users []*dbr.User

	err := d.db.Transaction(func(tx *gorm.DB) error {
		// Get the community members
		var members []*dbr.CommunityMember
		err := tx.Where("community_id = ?", communityID).
			Offset(offset).
			Limit(limit).
			Order("joined_at DESC").
			Find(&members).Error
		if err != nil {
			return err
		}

		// Extract user IDs
		var userIDs []uint
		for _, member := range members {
			userIDs = append(userIDs, member.UserID)
		}

		// Fetch user details
		if len(userIDs) > 0 {
			err = tx.Where("id IN ?", userIDs).Find(&users).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *communityDAO) AddMember(userID, communityID uint) error {
	return d.db.Create(&ms.CommunityMember{UserID: userID, CommunityID: communityID}).Error
}

func (d *communityDAO) RemoveMember(userID, communityID int64) error {
	return d.db.Where("user_id = ? AND community_id = ?", userID, communityID).Delete(&ms.CommunityMember{}).Error
}

func (d *communityDAO) ListMembers(communityID int64, offset, limit int) ([]*ms.User, error) {
	var users []*ms.User
	err := d.db.Joins("JOIN community_members ON users.id = community_members.user_id").
		Where("community_members.community_id = ?", communityID).
		Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}
