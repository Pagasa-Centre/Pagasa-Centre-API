package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

func EntityToDomainApprovals(entities []*entity.Approval) []*domain.Approval {
	if entities == nil {
		return nil
	}

	var approvals []*domain.Approval

	for _, ent := range entities {
		if ent == nil {
			continue
		}

		var approverID *string
		if ent.ApproverID.Valid {
			approverID = &ent.ApproverID.String
		}

		var approvalType string
		if ent.Type.Valid {
			approvalType = ent.Type.String
		}

		approval := &domain.Approval{
			ID:            ent.ID,
			RequesterID:   ent.RequesterID,
			ApproverID:    approverID,
			RequestedRole: ent.RequestedRole,
			Type:          approvalType,
			Status:        ent.Status,
		}

		approvals = append(approvals, approval)
	}

	return approvals
}
