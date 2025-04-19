package domain

type Approval struct {
	RequesterID   string
	ApproverID    string
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
	MinistryApplication = "Ministry Application"
)
