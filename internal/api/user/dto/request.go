package dto

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (cur CreateUserRequest) Validate() error {
	return validator.New().Struct(cur)
}
