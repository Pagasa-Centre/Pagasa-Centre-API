package mapper

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/domain"
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

		var updatedBy *string
		if ent.UpdatedBy.Valid {
			updatedBy = &ent.UpdatedBy.String
		}

		var reason string
		if ent.Reason.Valid {
			reason = ent.Reason.String
		}

		approval := &domain.Approval{
			ID:            ent.ID,
			RequesterID:   ent.RequesterID,
			UpdatedBy:     updatedBy,
			RequestedRole: ent.RequestedRole,
			Reason:        reason,
			Type:          ent.Type,
			Status:        ent.Status,
		}

		approvals = append(approvals, approval)
	}

	return approvals
}
