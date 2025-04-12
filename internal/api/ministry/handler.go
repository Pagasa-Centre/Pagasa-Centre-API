package ministry

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

type MinistryHandler interface {
	All() http.HandlerFunc
	Apply() http.HandlerFunc
}
type handler struct {
	logger          *zap.Logger
	MinistryService ministry.MinistryService
}

func NewMinistryHandler(
	logger *zap.Logger,
	ministryService ministry.MinistryService,
) MinistryHandler {
	return &handler{
		logger:          logger,
		MinistryService: ministryService,
	}
}

func (mh *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ministries, err := mh.MinistryService.All(ctx)
		if err != nil {
			mh.logger.Sugar().Infow("Failed to get all ministries", "error", err)
			render.Json(w, http.StatusInternalServerError, dto.ToErrorMinistriesResponse("Failed to fetch ministries"))

			return
		}

		resp := dto.ToGetAllMinistriesResponse(ministries, "Successfully fetched ministries")
		render.Json(w, http.StatusOK, resp)
	}
}

func (mh *handler) Apply() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.MinistryApplicationRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			mh.logger.Sugar().Errorw("failed to decode and validate login request body", "error", err)
			render.Json(w, http.StatusBadRequest,
				dto.ToMinistryApplicationResponse(
					"",
				),
			)

			return
		}

		// Get the user ID from the context
		userID, err := context.GetUserIDString(ctx)
		if err != nil {
			mh.logger.Sugar().Errorw("user ID not found in session", "error", err)
			render.Json(
				w,
				http.StatusUnauthorized,
				dto.ToMinistryApplicationResponse(
					"unauthorized",
				),
			)

			return
		}

		//1. Fetch Ministry Leader details(phone number & userid) via ministryID

		//2. Fetch User details(name, email and number)

		//3. Create New Approval (Type,requester_id, approver_id,state)

		//4. Construct and send Message to notify ministry leader that an application has been made

		//5. Let applicant know that their application has been sent.
	}
}
