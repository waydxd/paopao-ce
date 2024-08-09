package jinzhu

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"gorm.io/gorm"
)

type communityDAO struct {
	db *gorm.DB
}

func NewCommunityDAO(db *gorm.DB) ms.CommunityModelService {
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

func (d *communityDAO) AddMember(userID, communityID int64) error {
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
