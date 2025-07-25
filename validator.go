package gopackage

import (
	"github.com/go-playground/validator/v10"
)

// Custom Validator

type CustomValidatorData struct {
	validator *validator.Validate
}

func CustomValidator() *CustomValidatorData {
	return &CustomValidatorData{validator: validator.New()}
}

func (cvd *CustomValidatorData) Validate(i any) error {
	return cvd.validator.Struct(i)
}
