package core

import (
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/dao/jinzhu/dbr"
)

type ReportDAO interface {
	CreateReport(report *ms.Report) (int64, error)
	GetReportByID(id int64) (*ms.Report, error)
	UpdateReportStatus(id int64, status ms.Status) error
	DeleteReport(id int64) error
	ListReport(page, pageSize int) ([]ms.Report, error)
	ListReportByStatus(status string, page, pageSize int) ([]dbr.Report, int64, error)
}
