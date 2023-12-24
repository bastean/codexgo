package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Username struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

const UsernameMinCharactersLength = "2"
const UsernameMaxCharactersLength = "20"

var InvalidUsernameValue = errors.InvalidValue{Dump: []errors.Error{{Field: "Username", Message: "Must be between " + UsernameMinCharactersLength + " to " + UsernameMaxCharactersLength + " characters and be alphanumeric only"}}}

func ensureIsValidUsername(username *Username) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(username)

	return
}

func NewUsername(username string) *Username {
	usernameVO := &Username{username}

	err := ensureIsValidUsername(usernameVO)

	if err != nil {
		panic(InvalidUsernameValue)
	}

	return usernameVO
}
