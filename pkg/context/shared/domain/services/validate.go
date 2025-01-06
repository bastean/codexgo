package services

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled()).Struct
)

func IsValueObjectInvalid(object any) bool {
	return validate(object) != nil
}
