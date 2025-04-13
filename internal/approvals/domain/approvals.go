package domain

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type Approval struct {
	RequesterID string
	ApproverID  string
	Type        string
	Status      string
}

// Approval Types
const (
	MinistryApplicationType = "Ministry Application"
)

// Approval Status'
const (
	Pending  = "PENDING"
	Approved = "APPROVED"
	Rejected = "REJECTED"
)

func ToApprovalEntity(approval *Approval) *entity.Approval {
	return &entity.Approval{
		RequesterID: approval.RequesterID,
		ApproverID:  approval.ApproverID,
		Type:        approval.Type,
		Status:      approval.Status,
	}
}
