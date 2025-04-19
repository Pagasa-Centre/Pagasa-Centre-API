package domain

type Approval struct {
	RequesterID   string
	ApproverID    *string
	RequestedRole string
	Type          string
	Status        string
}

const (
	// Approval Status'
	Pending  = "PENDING"
	Approved = "APPROVED"
	Rejected = "REJECTED"

	// Approval Types
	MinistryApplication              = "Ministry Application"
	LeaderStatusConfirmation         = "Leader Status Confirmation"
	PrimaryStatusConfirmation        = "Primary Status Confirmation"
	PastorStatusConfirmation         = "Pastor Status Confirmation"
	MinistryLeaderStatusConfirmation = "Ministry Leader Status Confirmation"
)
