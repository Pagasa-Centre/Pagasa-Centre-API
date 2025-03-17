package domain

import "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user/dto"

type User struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
}

func CreateUserRequestToUserDomain(req dto.CreateUserRequest) (user User) {
	return User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Email:     req.Email,
	}
}
