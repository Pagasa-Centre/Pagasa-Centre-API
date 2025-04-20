package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"
	userDto "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	userDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

func ToGetAllApprovalsResponse(result *approvals.GetAllResult, message string) dto.GetAllApprovalsResponse {
	if result == nil {
		return dto.GetAllApprovalsResponse{
			Message: message,
		}
	}

	// Build a map for quick lookup of user details by ID
	userMap := make(map[string]*userDomain.User)

	for _, user := range result.Users {
		if user != nil {
			userMap[user.ID] = user
		}
	}

	var app []dto.Approval

	for _, a := range result.Approvals {
		if a == nil {
			continue
		}

		user := userMap[a.RequesterID]

		var requesterDetails userDto.UserDetails

		if user != nil {
			requesterDetails = userDto.UserDetails{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			}
		}

		app = append(app, dto.Approval{
			ID:               a.ID,
			Type:             a.Type,
			RequestedRole:    a.RequestedRole,
			Status:           a.Status,
			RequesterDetails: requesterDetails,
		})
	}

	return dto.GetAllApprovalsResponse{
		Approvals: app,
		Message:   message,
	}
}

func ToUpdateApprovalStatusResponse(message string) *dto.UpdateApprovalStatusResponse {
	return &dto.UpdateApprovalStatusResponse{
		Message: message,
	}
}
