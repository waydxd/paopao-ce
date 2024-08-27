package dbr

import (
	"gorm.io/gorm"
	"time"
)

type ReportType string

const (
	tweet   ReportType = "tweet"
	comment ReportType = "comment"
	// Add other types as needed
)

type Status string

const (
	Pending  Status = "pending"
	Reviewed Status = "reviewed"
	Resolved Status = "resolved"
)

type Report struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	PostID     int64     `gorm:"index" json:"post_id"`
	CommentID  int64     `gorm:"index" json:"comment_id"`
	Reason     string    `gorm:"type:varchar(255);not null" json:"reason"`
	ReporterID int64     `gorm:"not null" json:"reporter_id"`
	Status     Status    `gorm:"type:enum('pending','reviewed','resolved');default:'pending'" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// Create a new report
func (r *ReportRepository) Create(report *Report) error {
	return r.db.Create(report).Error
}

// Get a report by ID
func (r *ReportRepository) GetByID(id int64) (*Report, error) {
	var report Report
	err := r.db.Preload("Post").Preload("Reporter").First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// Update a report
func (r *ReportRepository) Update(id int64, status Status) error {
	return r.db.Model(&Report{}).Where("id = ?", id).Update("status", status).Error
}

// Delete a report
func (r *ReportRepository) Delete(id int64) error {
	return r.db.Delete(&Report{}, id).Error
}

// List all reports
func (r *ReportRepository) List(page, pageSize int) ([]Report, int64, error) {
	var reports []Report
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&Report{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Post").Preload("Reporter").
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&reports).Error

	if err != nil {
		return nil, 0, err
	}

	return reports, total, nil
}

// ListByStatus gets reports by status
func (r *ReportRepository) ListByStatus(status string, page, pageSize int) ([]Report, int64, error) {
	var reports []Report
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&Report{}).Where("status = ?", status).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Post").Preload("Reporter").
		Where("status = ?", status).
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&reports).Error

	if err != nil {
		return nil, 0, err
	}

	return reports, total, nil
}
