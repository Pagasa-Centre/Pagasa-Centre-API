package mapper

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/auth/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/domain"
	userdomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

func MapAuthResultToDTO(model *domain.AuthResult, message string) *dto.AuthResponse {
	if model == nil {
		return &dto.AuthResponse{
			Message: message,
		}
	}

	return &dto.AuthResponse{
		Token:   &model.AccessToken,
		User:    MapDomainUserToDtoUser(model.User, model.Roles),
		Message: message,
	}
}

func MapDomainUserToDtoUser(model *userdomain.User, roles []string) *dto.UserDetails {
	return &dto.UserDetails{
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		Email:        model.Email,
		Birthday:     model.Birthday.Format("2006-01-02"),
		PhoneNumber:  model.PhoneNumber,
		OutreachID:   model.OutreachID,
		CellLeaderID: model.CellLeaderID,
		Roles:        roles,
	}
}
