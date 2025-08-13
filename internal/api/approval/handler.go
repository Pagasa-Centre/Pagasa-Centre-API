package approval

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approval/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approval/dto/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type ApprovalHandler interface {
	All() http.HandlerFunc
	UpdateApprovalStatus() http.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	approvalService approval.Service
}

func NewApprovalHandler(
	logger *zap.Logger,
	approvalService approval.Service,
) ApprovalHandler {
	return &handler{
		logger:          logger,
		approvalService: approvalService,
	}
}

const (
	InternalServerErrorMsg = "Internal server error. Please try again later."
)

func (h *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := h.approvalService.GetAll(ctx)
		if err != nil {
			h.logger.Error("Failed to get all approvals", zap.Error(err))
			render.Json(
				w,
				http.StatusInternalServerError,
				mapper.ToGetAllApprovalsResponse(
					nil,
					InternalServerErrorMsg,
				))
		}

		render.Json(w, http.StatusOK, mapper.ToGetAllApprovalsResponse(
			result,
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
				mapper.ToUpdateApprovalStatusResponse(
					"approval id is required",
				),
			)

			return
		}

		var req dto.UpdateApprovalStatusRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			h.logger.Sugar().Errorw("failed to decode and validate update approval status body", "error", err)
			render.Json(w, http.StatusBadRequest,
				mapper.ToUpdateApprovalStatusResponse(
					"Please select a valid approval status",
				),
			)

			return
		}

		err := h.approvalService.UpdateApprovalStatus(ctx, approvalID, req.Status)
		if err != nil {
			h.logger.Sugar().Errorw("Failed to update approval status", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(w, statusCode,
				mapper.ToUpdateApprovalStatusResponse(
					errMsg,
				),
			)

			return
		}

		render.Json(w, http.StatusOK,
			mapper.ToUpdateApprovalStatusResponse(
				"Successfully updated approval status",
			),
		)
	}
}

func mapErrorsToStatusCodeAndUserFriendlyMessages(err error) (int, string) {
	switch {
	case errors.Is(err, approval.ErrNoPermission):
		return http.StatusForbidden, "You do not have permission to approve or reject this application."
	default:
		return http.StatusInternalServerError, InternalServerErrorMsg
	}
}
