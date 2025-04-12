package dto

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type UpdateUserDetailsResponse struct {
	Message string       `json:"message"`
	User    *UserDetails `json:"user,omitempty"`
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
	var userDetails *UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &RegisterResponse{
		Token:   token,
		User:    userDetails,
		Message: message,
	}
}

func ToLoginResponse(user *entity.User, token string, message string) *LoginResponse {
	var userDetails *UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &LoginResponse{
		Token:   &token,
		User:    userDetails,
		Message: message,
	}
}

func ToUpdateUserDetailsResponse(user *entity.User, message string) *UpdateUserDetailsResponse {
	var userDetails *UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &UpdateUserDetailsResponse{
		Message: message,
		User:    userDetails,
	}
}

func extractUserDetails(user *entity.User) *UserDetails {
	var birthday, phone, outreachID string

	var cellLeaderID *string

	if user.Birthday.Valid {
		birthday = user.Birthday.Time.Format("2006-01-02")
	}

	if user.Phone.Valid {
		phone = user.Phone.String
	}

	if user.OutreachID.Valid {
		outreachID = user.OutreachID.String
	}

	if user.CellLeaderID.Valid {
		cellLeaderID = &user.CellLeaderID.String
	}

	return &UserDetails{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Birthday:     birthday,
		PhoneNumber:  phone,
		OutreachID:   outreachID,
		CellLeaderID: cellLeaderID,
	}
}
