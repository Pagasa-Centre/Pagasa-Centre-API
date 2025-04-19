package mappers

import (
	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

func ToApprovalEntity(approval *domain.Approval) *entity.Approval {
	return &entity.Approval{
		RequesterID:   approval.RequesterID,
		ApproverID:    null.StringFromPtr(approval.ApproverID),
		RequestedRole: approval.RequestedRole,
		Type:          null.StringFrom(approval.Type),
		Status:        approval.Status,
	}
}
