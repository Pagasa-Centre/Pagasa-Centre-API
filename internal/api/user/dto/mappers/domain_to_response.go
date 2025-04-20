package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

func ToDeleteUserResponse(message string) *dto.DeleteUserResponse {
	return &dto.DeleteUserResponse{
		Message: message,
	}
}

func ToRegisterResponse(registerResult *user.RegisterResult, message string) *dto.RegisterResponse {
	if registerResult == nil {
		return &dto.RegisterResponse{
			Message: message,
		}
	}

	return &dto.RegisterResponse{
		Token: &registerResult.Token,
		User: &dto.UserDetails{
			FirstName:    registerResult.User.FirstName,
			LastName:     registerResult.User.LastName,
			Email:        registerResult.User.Email,
			Birthday:     registerResult.User.Birthday.Format("2006-01-02"),
			PhoneNumber:  registerResult.User.PhoneNumber,
			OutreachID:   registerResult.User.OutreachID,
			CellLeaderID: registerResult.User.CellLeaderID,
			Roles:        registerResult.Roles,
		},
		Message: message,
	}
}

func ToLoginResponse(loginResult *user.AuthResult, message string) *dto.LoginResponse {
	if loginResult == nil {
		return &dto.LoginResponse{
			Message: message,
		}
	}

	return &dto.LoginResponse{
		Token: &loginResult.Token,
		User: &dto.UserDetails{
			FirstName:    loginResult.User.FirstName,
			LastName:     loginResult.User.LastName,
			Email:        loginResult.User.Email,
			Birthday:     loginResult.User.Birthday.Format("2006-01-02"),
			PhoneNumber:  loginResult.User.PhoneNumber,
			OutreachID:   loginResult.User.OutreachID,
			CellLeaderID: loginResult.User.CellLeaderID,
			Roles:        loginResult.Roles,
		},
		Message: message,
	}
}

func ToUpdateUserDetailsResponse(user *domain.User, message string) *dto.UpdateUserDetailsResponse {
	if user == nil {
		return &dto.UpdateUserDetailsResponse{
			Message: message,
		}
	}

	return &dto.UpdateUserDetailsResponse{
		Message: message,
		User: &dto.UserDetails{
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
