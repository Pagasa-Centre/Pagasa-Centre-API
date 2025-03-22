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

type UpdateDetailsRequest struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty" validate:"omitempty,email"`
	Password     string `json:"password,omitempty"`
	Birthday     string `json:"birthday,omitempty"` // Format: "2006-01-02"
	PhoneNumber  string `json:"phone_number,omitempty"`
	CellLeaderID *int   `json:"cell_leader_id,omitempty"`
	OutreachID   int    `json:"outreach_id,omitempty"`
}

func (rr RegisterRequest) Validate() error {
	return validator.New().Struct(rr)
}

func (lr LoginRequest) Validate() error {
	return validator.New().Struct(lr)
}

func (udr UpdateDetailsRequest) Validate() error {
	return validator.New().Struct(udr)
}
