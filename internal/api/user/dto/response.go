package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type UpdateUserDetailsResponse struct {
	Message string
	User    *UserDetails
}

type UserDetails struct {
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Email        string  `json:"email,omitempty" validate:"omitempty,email"`
	Birthday     string  `json:"birthday,omitempty"`
	PhoneNumber  string  `json:"phone_number,omitempty"`
	CellLeaderID *string `json:"cell_leader_id,omitempty"`
	OutreachID   string  `json:"outreach_id,omitempty"`
}

type RegisterResponse struct {
	Token   *string      `json:"token,omitempty"`
	User    *UserDetails `json:"user,omitempty"`
	Message string       `json:"message"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	Token   *string      `json:"token,omitempty"`
	User    *UserDetails `json:"user,omitempty"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func ToDeleteUserResponse(message string) *DeleteUserResponse {
	return &DeleteUserResponse{
		Message: message,
	}
}

func ToRegisterResponse(user *entity.User, token *string, message string) *RegisterResponse {
	var birthday string
	if user.Birthday.Valid {
		birthday = user.Birthday.Time.Format("2006-01-02") // Format date to YYYY-MM-DD
	}

	var phone string
	if user.Phone.Valid {
		phone = user.Phone.String
	}

	var outreachID string
	if user.OutreachID.Valid {
		outreachID = user.OutreachID.String
	}

	var cellLeaderID *string
	if user.CellLeaderID.Valid {
		cellLeaderID = &user.CellLeaderID.String
	}

	return &RegisterResponse{
		Token: token,
		User: &UserDetails{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Birthday:     birthday,
			PhoneNumber:  phone,
			OutreachID:   outreachID,
			CellLeaderID: cellLeaderID,
		},
		Message: message,
	}
}

func ToLoginResponse(user *entity.User, token string, message string) *LoginResponse {
	var birthday string
	if user.Birthday.Valid {
		// Format the birthday as YYYY-MM-DD
		birthday = user.Birthday.Time.Format("2006-01-02")
	}

	var phone string
	if user.Phone.Valid {
		phone = user.Phone.String
	}

	var cellLeaderID *string

	if user.CellLeaderID.Valid {
		// Convert null.Int to int and assign its address.
		v := user.CellLeaderID.String
		cellLeaderID = &v
	}

	var outreachID string
	if user.OutreachID.Valid {
		outreachID = user.OutreachID.String
	}

	return &LoginResponse{
		Token: &token,
		User: &UserDetails{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Birthday:     birthday,
			PhoneNumber:  phone,
			CellLeaderID: cellLeaderID,
			OutreachID:   outreachID,
		},
		Message: message,
	}
}

func ToUpdateUserDetailsResponse(user *entity.User, message string) *UpdateUserDetailsResponse {
	var birthday string
	if user.Birthday.Valid {
		// Format the birthday as YYYY-MM-DD
		birthday = user.Birthday.Time.Format("2006-01-02")
	}

	var phone string
	if user.Phone.Valid {
		phone = user.Phone.String
	}

	var cellLeaderID *string

	if user.CellLeaderID.Valid {
		// Convert null.Int to int and assign its address.
		v := user.CellLeaderID.String
		cellLeaderID = &v
	}

	var outreachID string
	if user.OutreachID.Valid {
		outreachID = user.OutreachID.String
	}

	return &UpdateUserDetailsResponse{
		Message: message,
		User: &UserDetails{
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Birthday:     birthday,
			PhoneNumber:  phone,
			CellLeaderID: cellLeaderID,
			OutreachID:   outreachID,
		},
	}
}
