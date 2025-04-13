package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Accepts either international format (+447...) or local UK format (07...) — 11 digits.
func E164OrLocalPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	// E.164 (e.g., +447958662569)
	isE164 := regexp.MustCompile(`^\+44\d{10}$`).MatchString

	// UK local (e.g., 07958662569)
	isUKLocal := regexp.MustCompile(`^07\d{9}$`).MatchString

	return isE164(phone) || isUKLocal(phone)
}

// Register the custom validator to be used globally
func RegisterCustomPhoneValidator(v *validator.Validate) {
	_ = v.RegisterValidation("e164orlocal", E164OrLocalPhone)
}
