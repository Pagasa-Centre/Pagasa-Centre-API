package dto

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	Birthday     string `json:"birthday" validate:"required"`
	OutreachID   int    `json:"outreach_id" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	CellLeaderID int    `json:"cell_leader_id,omitempty"`
	IsLeader     bool   `json:"is_leader" `
	IsPrimary    bool   `json:"is_primary" `
	IsPastor     bool   `json:"is_pastor"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (rr RegisterRequest) Validate() error {
	return validator.New().Struct(rr)
}

func (lr LoginRequest) Validate() error {
	return validator.New().Struct(lr)
}
