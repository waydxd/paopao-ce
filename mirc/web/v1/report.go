package v1

import (
	. "github.com/alimy/mir/v4"
	. "github.com/alimy/mir/v4/engine"
	"github.com/waydxd/paopao-ce/internal/model/web"
)

func init() {
	Entry[Report]()
}

type Report struct {
	Chain `mir:"-"`
	Group `mir:"v1"`
	// 檢舉 report
	SendReport func(Post, web.ReportReq) web.ReportResp `mir:"/report"`

	DeleteReport func(Delete, web.DeleteReportReq) web.ReportResp `mir:"/report"`
	// List reported posts 獲取被檢舉的帖子
	ListReported func(Get) web.ListReportResp `mir:"/report/list"`
}
