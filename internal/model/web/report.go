package web

import "github.com/waydxd/paopao-ce/internal/core/ms"

type ReportType string

const (
	Tweet   ReportType = "tweet"
	Comment ReportType = "comment"
	// Add other types as needed
)

type ReportReq struct {
	PostId    int64  `json:"post_id"`
	CommentID int64  `json:"comment_id"`
	Reason    string `json:"reason" binding:"required"`
	UserId    int64  `json:"user_id" binding:"required"`
}

type ReportResp struct {
	ReportId string `json:"report_id"`
}

type ListReportResp struct {
	Posts []ms.Report `json:"reports"`
}
type DeleteReportReq struct {
	ID int64 `json:"id" binding:"required"`
}
type PatchReportReq struct {
	ReportId int64     `form:"report_id" binding:"required"`
	Status   ms.Status `form:"status" binding:"required"`
}
