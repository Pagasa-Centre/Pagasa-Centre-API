package mapper

import (
	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

func ToApprovalEntity(approval *domain.Approval) *entity.Approval {
	return &entity.Approval{
		RequesterID:   approval.RequesterID,
		RequestedRole: approval.RequestedRole,
		Reason:        null.StringFrom(approval.Reason),
		UpdatedBy:     null.StringFromPtr(approval.UpdatedBy),
		Type:          approval.Type,
		Status:        approval.Status,
		MinistryID:    null.StringFromPtr(approval.MinistryID),
	}
}
