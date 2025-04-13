package approvals

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type ApprovalHandler interface {
	GetAll() http.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	approvalService approvals.ApprovalService
}

func NewApprovalHandler(
	logger *zap.Logger,
	approvalService approvals.ApprovalService,
) ApprovalHandler {
	return &handler{
		logger:          logger,
		approvalService: approvalService,
	}
}

func (h *handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		a, err := h.approvalService.GetAllApprovals(ctx)
		if err != nil {
			h.logger.Error("Failed to get all approvals", zap.Error(err))
			render.Json(
				w,
				http.StatusInternalServerError,
				dto.ToGetAllApprovalsResponse(
					nil,
					"Failed to get all approvals",
				))
		}

		render.Json(w, http.StatusOK, dto.ToGetAllApprovalsResponse(
			&a,
			"Successfully got all approvals",
		))
	}
}
