package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
)

type Approval struct {
	ID               string          `json:"id"`
	Type             string          `json:"type"`
	RequestedRole    string          `json:"requested_role"`
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
