package domain

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
