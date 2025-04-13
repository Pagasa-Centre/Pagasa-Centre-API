package dto

import (
	"github.com/go-playground/validator/v10"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/validators"
)

type RegisterRequest struct {
	FirstName        string  `json:"first_name" validate:"required"`
	LastName         string  `json:"last_name" validate:"required"`
	Email            string  `json:"email" validate:"required,email"`
	Password         string  `json:"password" validate:"required"`
	Birthday         string  `json:"birthday" validate:"required"`
	OutreachID       string  `json:"outreach_id" validate:"required"`
	PhoneNumber      string  `json:"phone_number" validate:"required,e164orlocal"`
	CellLeaderID     *string `json:"cell_leader_id,omitempty"`
	MinistryID       *string `json:"ministry_id,omitempty"`
	IsLeader         bool    `json:"is_leader"`
	IsPrimary        bool    `json:"is_primary"`
	IsPastor         bool    `json:"is_pastor"`
	IsMinistryLeader bool    `json:"is_ministry_leader"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateDetailsRequest struct {
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Email        string  `json:"email,omitempty" validate:"omitempty,email"`
	Password     string  `json:"password,omitempty"`
	Birthday     string  `json:"birthday,omitempty"`
	PhoneNumber  string  `json:"phone_number,omitempty" validate:"omitempty,e164orlocal"`
	CellLeaderID *string `json:"cell_leader_id,omitempty"`
	OutreachID   string  `json:"outreach_id,omitempty"`
}

func (rr RegisterRequest) Validate() error {
	v := validator.New()
	validators.RegisterCustomPhoneValidator(v)

	return v.Struct(rr)
}

func (lr LoginRequest) Validate() error {
	return validator.New().Struct(lr)
}

func (udr UpdateDetailsRequest) Validate() error {
	v := validator.New()
	validators.RegisterCustomPhoneValidator(v)

	return v.Struct(udr)
}
