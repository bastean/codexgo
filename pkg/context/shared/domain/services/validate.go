package services

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled()).Struct
)

func IsValueObjectInvalid(valueObj any) bool {
	return validate(valueObj) != nil
}
