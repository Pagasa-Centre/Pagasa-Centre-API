package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
)

type Approval struct {
	ID               string          `json:"id"`
	Type             string          `json:"type"`
	Status           string          `json:"status"`
	RequesterDetails dto.UserDetails `json:"requester_details"`
}

type UpdateApprovalStatusResponse struct {
	Message string `json:"message"`
}

type GetAllApprovalsResponse struct {
	Approvals *[]Approval `json:"approvals"`
	Message   string      `json:"message"`
}

func ToGetAllApprovalsResponse(approvals *[]Approval, message string) *GetAllApprovalsResponse {
	return &GetAllApprovalsResponse{
		Approvals: approvals,
		Message:   message,
	}
}

func ToUpdateApprovalStatusResponse(message string) *UpdateApprovalStatusResponse {
	return &UpdateApprovalStatusResponse{
		Message: message,
	}
}
