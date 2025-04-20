package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
)

func UserEntityToUserDomain(user *entity.User) *domain.User {
	var phoneNumber string
	if user.Phone.Valid {
		phoneNumber = user.Phone.String
	}

	var cellLeaderID *string
	if user.CellLeaderID.Valid {
		cellLeaderID = &user.CellLeaderID.String
	}

	return &domain.User{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Birthday:       user.Birthday.Time,
		PhoneNumber:    phoneNumber,
		CellLeaderID:   cellLeaderID,
		OutreachID:     user.OutreachID.String,
	}
}
