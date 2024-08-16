// community.go

package dbr

import (
	"gorm.io/gorm"
	"time"
)

// Model definitions

type CommunityMember struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CommunityID uint      `json:"community_id"`
	UserID      uint      `json:"user_id"`
	Role        string    `json:"role"`
	JoinedAt    time.Time `json:"joined_at"`
}

type Community struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex"`
	Slug        string    `json:"slug" gorm:"uniqueIndex"`
	Description string    `json:"description"`
	AvatarURL   string    `json:"avatar_url"`
	BannerURL   string    `json:"banner_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type CommunityTag struct {
	// TODO
}

// CRUD Op
func (c *Community) Create(db *gorm.DB) (*Community, error) {
	err := db.Create(&c).Error
	return c, err
}

func (c *Community) Get(db *gorm.DB) (*Community, error) {
	err := db.First(&c, c.ID).Error
	return c, err
}

func (c *Community) Update(db *gorm.DB) error {
	return db.Save(c).Error
}

func (c *Community) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

func (c *Community) List(db *gorm.DB, offset, limit int) ([]*Community, int64, error) {
	var communities []*Community
	var count int64
	err := db.Model(&Community{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset(offset).Limit(limit).Find(&communities).Error
	return communities, count, err
}

func (cm *CommunityMember) Create(db *gorm.DB) (*CommunityMember, error) {
	err := db.Create(&cm).Error
	return cm, err
}

// Service methods (if needed)
// type CommunityService struct {
// 	DB *gorm.DB
// }

// func (s *CommunityService) CreateCommunity(name, description string, creatorID uint) (*Community, error) {
// 	community := &Community{
// 		Name:        name,
// 		Description: description,
// 	}
// 	_, err := community.Create(s.DB)
// 	if err != nil {
// 		return nil, err
// 	}

// 	member := &CommunityMember{
// 		CommunityID: community.ID,
// 		UserID:      creatorID,
// 		Role:        "admin",
// 		JoinedAt:    time.Now(),
// 	}
// 	_, err = member.Create(s.DB)
// 	if err != nil {
// 		// You might want to delete the community if member creation fails
// 		return nil, err
// 	}

// 	return community, nil
// }
