// community.go

package dbr // Assuming this is your package name

import (
	"time"

	"gorm.io/gorm"
)

// Model definitions
type Community struct {
	Model       // Assuming you have a base Model struct
	ID          uint
	Name        string `gorm:"uniqueIndex;not null"`
	Slug        string `gorm:"uniqueIndex;not null"`
	Description string
	AvatarURL   string
	BannerURL   string
}

type CommunityMember struct {
	Model
	CommunityID uint
	UserID      uint
	Role        string
	JoinedAt    time.Time
}

// Repository methods

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
type CommunityService struct {
	DB *gorm.DB
}

func (s *CommunityService) CreateCommunity(name, description string, creatorID uint) (*Community, error) {
	community := &Community{
		Name:        name,
		Description: description,
	}
	_, err := community.Create(s.DB)
	if err != nil {
		return nil, err
	}

	member := &CommunityMember{
		CommunityID: community.ID,
		UserID:      creatorID,
		Role:        "admin",
		JoinedAt:    time.Now(),
	}
	_, err = member.Create(s.DB)
	if err != nil {
		// You might want to delete the community if member creation fails
		return nil, err
	}

	return community, nil
}
