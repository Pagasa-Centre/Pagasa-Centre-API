package mappers

import "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals/dto"

func ToGetAllApprovalsResponse(approvals *[]dto.Approval, message string) *dto.GetAllApprovalsResponse {
	return &dto.GetAllApprovalsResponse{
		Approvals: approvals,
		Message:   message,
	}
}

func ToUpdateApprovalStatusResponse(message string) *dto.UpdateApprovalStatusResponse {
	return &dto.UpdateApprovalStatusResponse{
		Message: message,
	}
}
