package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	v             = validator.New()
	validationMap = map[string]func(validator.FieldLevel) bool{
		"validCPF": validCPF,
	}
)

func RegisterValidation() error {
	return registerValidationMap(validationMap)
}

func Validate(entity interface{}) error {
	return v.Struct(entity)
}

func registerValidationMap(m map[string]func(validator.FieldLevel) bool) error {
	for key, value := range m {
		if err := v.RegisterValidation(key, value); err != nil {
			return err
		}
	}
	return nil
}

func validCPF(fl validator.FieldLevel) bool {
	return IsValidCPF(fl.Field().String())
}
