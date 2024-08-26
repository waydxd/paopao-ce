package web

import (
	"net/http"
	"strconv"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/waydxd/paopao-ce/auto/api/v1"
	"github.com/waydxd/paopao-ce/internal/core/ms"
	"github.com/waydxd/paopao-ce/internal/model/web"
	"github.com/waydxd/paopao-ce/internal/servants/base"
	"github.com/waydxd/paopao-ce/internal/servants/chain"
)

type ReportSrv struct {
	api.UnimplementedReportServant
	*base.DaoServant
}

func NewReportSrv(s *base.DaoServant) *ReportSrv {
	return &ReportSrv{
		DaoServant: s,
	}
}
func (s *ReportSrv) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}
func (s *ReportSrv) ListReported() (*web.ListReportResp, mir.Error) {
	posts, err := s.Ds.ListReport(1, 100) // Adjust page and pageSize as needed
	if err != nil {
		return nil, mir.NewError(http.StatusInternalServerError, web.ErrFetchReport)
	}
	return &web.ListReportResp{Posts: posts}, nil
}

func (s *ReportSrv) DeleteReport(req *web.DeleteReportReq) (*web.ReportResp, mir.Error) {
	err := s.Ds.DeleteReport(req.ID)
	if err != nil {
		return nil, mir.NewError(http.StatusInternalServerError, web.ErrDeleteReport)
	}

	return &web.ReportResp{ReportId: strconv.FormatInt(req.ID, 10)}, nil
}

func (s *ReportSrv) SendReport(req *web.ReportReq) (*web.ReportResp, mir.Error) {
	report := &ms.Report{
		//ReportType: dbr.ReportType(req.ReportType),
		PostID:     req.PostId,
		CommentID:  req.CommentID,
		Reason:     req.Reason,
		ReporterID: req.UserId,
		Status:     "pending",
	}
	id, err := s.Ds.CreateReport(report)
	if err != nil {
		return nil, mir.NewError(http.StatusInternalServerError, web.ErrSendReport)
	}

	return &web.ReportResp{ReportId: strconv.FormatInt(id, 10)}, nil
}
