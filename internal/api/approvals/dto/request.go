package dto

import "github.com/go-playground/validator/v10"

type UpdateApprovalStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

func (u UpdateApprovalStatusRequest) Validate() error {
	return validator.New().Struct(u)
}
