package jinzhu

import (
	"github.com/sirupsen/logrus"
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type communityDAO struct {
	db *gorm.DB
}

func (d *communityDAO) GetCommunityPost(communityID uint, offset int, limit int) ([]*ms.Post, int64, error) {
	predicates := dbr.Predicates{
		"community_id = ?": []any{communityID},
		"ORDER":            []any{"created_on DESC"},
	}

	// Fetch posts
	posts, err := (&dbr.Post{}).Fetch(d.db, predicates, offset, limit)
	if err != nil {
		logrus.Debugf("communityManageDAO.GetCommunityPost fetch err: %v", err)
		return nil, 0, err
	}

	// Get total count
	total, err := (&dbr.Post{}).CountBy(d.db, predicates)
	if err != nil {
		logrus.Debugf("communityManageDAO.GetCommunityPost count err: %v", err)
		return nil, 0, err
	}

	return posts, total, nil
}

type communityManageDAO struct {
	db *gorm.DB
}

// ListCommunityMembers implements core.CommunityService.

func newCommunityService(db *gorm.DB) core.CommunityService {
	return &communityDAO{db: db}
}

func newCommunityManageService(db *gorm.DB) core.CommunityManageService {
	return &communityManageDAO{db: db}
}

func (d *communityManageDAO) CreateCommunity(community *ms.Community) error {
	return d.db.Create(community).Error
}

func (d *communityManageDAO) RemoveCommunity(communityID uint) error {
	return d.db.Where("id = ?", communityID).Delete(&ms.Community{}).Error
}

func (d *communityDAO) GetCommunity(communityID uint) (*ms.Community, error) {
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

func (d *communityDAO) ListCommunityMembers(communityID uint, offset int, limit int) ([]*dbr.User, error) {
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

func (d *communityManageDAO) AddMember(userID, communityID uint) error {
	return d.db.Create(&ms.CommunityMember{UserID: userID, CommunityID: communityID}).Error
}

func (d *communityManageDAO) RemoveMember(userID, communityID uint) error {
	return d.db.Where("user_id = ? AND community_id = ?", userID, communityID).Delete(&ms.CommunityMember{}).Error
}

func (d *communityManageDAO) JoinCommunity(userID, communityID uint) error {
	return d.db.Create(&ms.CommunityMember{UserID: userID, CommunityID: communityID}).Error
}

func (d *communityManageDAO) LeaveCommunity(userID, communityID uint) error {
	return d.db.Where("user_id = ? AND community_id = ?", userID, communityID).Delete(&ms.CommunityMember{}).Error
}

// func (d *communityManageDAO) ListMembers(communityID uint, offset, limit int) ([]*ms.User, error) {
// 	var users []*ms.User
// 	err := d.db.Joins("JOIN community_members ON users.id = community_members.user_id").
// 		Where("community_members.community_id = ?", communityID).
// 		Offset(offset).Limit(limit).Find(&users).Error
// 	return users, err
// }
