package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type UpdateDetailsResponse struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty" validate:"omitempty,email"`
	Birthday     string `json:"birthday,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	CellLeaderID *int   `json:"cell_leader_id,omitempty"`
	OutreachID   int    `json:"outreach_id,omitempty"`
}

func ToResponse(user *entity.User) (update UpdateDetailsResponse) {
	var birthday string
	if user.Birthday.Valid {
		// Format the birthday as YYYY-MM-DD
		birthday = user.Birthday.Time.Format("2006-01-02")
	}

	var phone string
	if user.Phone.Valid {
		phone = user.Phone.String
	}

	var cellLeaderID *int

	if user.CellLeaderID.Valid {
		// Convert null.Int to int and assign its address.
		v := user.CellLeaderID.Int
		cellLeaderID = &v
	}

	var outreachID int
	if user.OutreachID.Valid {
		outreachID = user.OutreachID.Int
	}

	return UpdateDetailsResponse{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Birthday:     birthday,
		PhoneNumber:  phone,
		CellLeaderID: cellLeaderID,
		OutreachID:   outreachID,
	}
}
