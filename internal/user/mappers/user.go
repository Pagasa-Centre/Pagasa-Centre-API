package mappers

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/domain"
	"github.com/volatiletech/null/v8"
)

func ToUserEntity(user domain.User) *entity.User {
	return &entity.User{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Birthday:       null.TimeFrom(user.Birthday),
		Phone:          null.StringFrom(user.PhoneNumber),
		CellLeaderID:   null.IntFromPtr(user.CellLeaderID),
		OutreachID:     null.IntFrom(user.OutreachID),
	}
}
