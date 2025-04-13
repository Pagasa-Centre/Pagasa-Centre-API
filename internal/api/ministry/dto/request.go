package dto

import "github.com/go-playground/validator/v10"

type MinistryApplicationRequest struct {
	MinistryID string `json:"ministry_id" validate:"required"`
	Reason     string `json:"reason" validate:"required"`
}

func (rr MinistryApplicationRequest) Validate() error {
	return validator.New().Struct(rr)
}
