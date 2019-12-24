package presenter

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	validator := &Validator{validator: validator.New()}

	return validator
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
