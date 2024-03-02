package valueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Username struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

const UsernameMinCharactersLength = "2"
const UsernameMaxCharactersLength = "20"

var InvalidUsernameValue = errors.InvalidValue{Message: "Username must be between " + UsernameMinCharactersLength + " to " + UsernameMaxCharactersLength + " characters and be alphanumeric only"}

func ensureIsValidUsername(username *Username) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(username)
}

func NewUsername(username string) *Username {
	username = strings.TrimSpace(username)

	usernameVO := &Username{username}

	err := ensureIsValidUsername(usernameVO)

	if err != nil {
		panic(InvalidUsernameValue)
	}

	return usernameVO
}
