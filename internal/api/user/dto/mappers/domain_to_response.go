package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

func ToDeleteUserResponse(message string) *dto.DeleteUserResponse {
	return &dto.DeleteUserResponse{
		Message: message,
	}
}

func ToRegisterResponse(user *entity.User, token *string, message string) *dto.RegisterResponse {
	var userDetails *dto.UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &dto.RegisterResponse{
		Token:   token,
		User:    userDetails,
		Message: message,
	}
}

func ToLoginResponse(user *entity.User, token string, message string) *dto.LoginResponse {
	var userDetails *dto.UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &dto.LoginResponse{
		Token:   &token,
		User:    userDetails,
		Message: message,
	}
}

func ToUpdateUserDetailsResponse(user *entity.User, message string) *dto.UpdateUserDetailsResponse {
	var userDetails *dto.UserDetails
	if user != nil {
		userDetails = extractUserDetails(user)
	}

	return &dto.UpdateUserDetailsResponse{
		Message: message,
		User:    userDetails,
	}
}

func extractUserDetails(user *entity.User) *dto.UserDetails {
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

	return &dto.UserDetails{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Birthday:     birthday,
		PhoneNumber:  phone,
		OutreachID:   outreachID,
		CellLeaderID: cellLeaderID,
	}
}
