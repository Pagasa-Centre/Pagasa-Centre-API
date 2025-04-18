package domain

import (
	"github.com/volatiletech/null/v8"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type Approval struct {
	RequesterID   string
	ApproverID    string
	RequestedRole string
	Type          string
	Status        string
}

// Approval Status'
const (
	Pending             = "PENDING"
	Approved            = "APPROVED"
	Rejected            = "REJECTED"
	MinistryApplication = "Ministry Application"
)

func ToApprovalEntity(approval *Approval) *entity.Approval {
	return &entity.Approval{
		RequesterID:   approval.RequesterID,
		ApproverID:    approval.ApproverID,
		RequestedRole: approval.RequestedRole,
		Type:          null.StringFrom(approval.Type),
		Status:        approval.Status,
	}
}
