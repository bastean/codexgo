package values

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New().Var
)

func IsNotValid(object any, validations ...string) bool {
	return validate(object, strings.Join(validations, ",")) != nil
}
