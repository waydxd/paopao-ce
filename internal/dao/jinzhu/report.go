package jinzhu

import (
	"github.com/waydxd/paopao-ce/internal/core"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type ReportDAO struct {
	db *gorm.DB
	r  dbr.ReportRepository
}

func NewReportDAO(db *gorm.DB) core.ReportDAO {
	return &ReportDAO{
		db: db,
		r:  *dbr.NewReportRepository(db),
	}
}

func (dao *ReportDAO) CreateReport(report *dbr.Report) (int64, error) {

	result := dao.db.Create(report)
	if result.Error != nil {
		return 0, result.Error
	}
	return report.ID, nil
}

func (dao *ReportDAO) GetReportByID(id int64) (*dbr.Report, error) {
	var report *dbr.Report
	report, err := dao.r.GetByID(id)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (dao *ReportDAO) UpdateReport(report *dbr.Report) error {
	return dao.r.Update(report)
}

func (dao *ReportDAO) DeleteReport(id int64) error {
	return dao.r.Delete(id)
}
func (dao *ReportDAO) ListReport(page, pageSize int) ([]ms.Report, error) {
	var reports []ms.Report
	offset := (page - 1) * pageSize

	err := dao.db.Order("created_at DESC").
		Offset(offset).Limit(pageSize).
		Find(&reports).Error

	if err != nil {
		return nil, err
	}

	// Assuming ms.IndexTweetList is a struct that can be populated from the reports
	// You might need to adjust this part based on the actual structure of ms.IndexTweetList
	// Populate tweetList from reports

	return reports, nil
}

func (dao *ReportDAO) ListReportByStatus(status string, page, pageSize int) ([]dbr.Report, int64, error) {
	return dao.r.ListByStatus(status, page, pageSize)
}
