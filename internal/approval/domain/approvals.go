package domain

type (
	Status string
	Type   string
)

const (
	// Approval Status'
	Pending  Status = "PENDING"
	Approved Status = "APPROVED"
	Rejected Status = "REJECTED"

	// Approval Types
	MinistryApplication              Type = "Ministry Application"
	LeaderStatusConfirmation         Type = "Leader Status Confirmation"
	PrimaryStatusConfirmation        Type = "Primary Status Confirmation"
	PastorStatusConfirmation         Type = "Pastor Status Confirmation"
	MinistryLeaderStatusConfirmation Type = "Ministry Leader Status Confirmation"
)

type Approval struct {
	ID            string
	RequesterID   string
	UpdatedBy     *string
	RequestedRole string
	Type          string
	Status        string
	Reason        string
	MinistryID    *string
}
