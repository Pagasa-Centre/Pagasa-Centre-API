package approvals

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type ApprovalHandler interface {
	All() http.HandlerFunc
	UpdateApprovalStatus() http.HandlerFunc
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

func (h *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		a, err := h.approvalService.GetAllApprovals(ctx)
		if err != nil {
			h.logger.Error("Failed to get all approvals", zap.Error(err))
			render.Json(
				w,
				http.StatusInternalServerError,
				mappers.ToGetAllApprovalsResponse(
					nil,
					"Failed to get all approvals",
				))
		}

		render.Json(w, http.StatusOK, mappers.ToGetAllApprovalsResponse(
			&a,
			"Successfully got all approvals",
		))
	}
}

func (h *handler) UpdateApprovalStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		approvalID := chi.URLParam(r, "id")
		if approvalID == "" {
			render.Json(w, http.StatusBadRequest,
				mappers.ToUpdateApprovalStatusResponse(
					"id is required",
				),
			)

			return
		}

		var req dto.UpdateApprovalStatusRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate update approval status body", "error", err)
			render.Json(w, http.StatusBadRequest,
				mappers.ToUpdateApprovalStatusResponse(
					"Please select a valid approval status",
				),
			)

			return
		}

		err := h.approvalService.UpdateApprovalStatus(ctx, approvalID, req.Status)
		if err != nil {
			h.logger.Sugar().Errorw("Failed to update approval status", "error", err)
			render.Json(w, http.StatusInternalServerError,
				mappers.ToUpdateApprovalStatusResponse(
					"something went wrong. please try again later",
				),
			)

			return
		}

		render.Json(w, http.StatusOK,
			mappers.ToUpdateApprovalStatusResponse(
				"Successfully updated approval status",
			),
		)
	}
}
