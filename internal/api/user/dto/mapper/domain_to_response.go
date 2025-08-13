package mapper

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

func ToUpdateUserDetailsResponse(user *domain.User, message string) *dto.UpdateUserDetailsResponse {
	if user == nil {
		return &dto.UpdateUserDetailsResponse{
			Message: message,
		}
	}

	return &dto.UpdateUserDetailsResponse{
		Message: message,
		User: &dto.UserDetails{ // todo: refactor into toUserDetailsDTO func
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			Birthday:     user.Birthday.Format("2006-01-02"),
			PhoneNumber:  user.PhoneNumber,
			OutreachID:   user.OutreachID,
			CellLeaderID: user.CellLeaderID,
		},
	}
}

func ToDeleteUserResponse(message string) *dto.DeleteUserResponse {
	return &dto.DeleteUserResponse{
		Message: message,
	}
}
