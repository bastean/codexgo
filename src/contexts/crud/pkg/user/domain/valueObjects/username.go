package valueObjects

import "github.com/go-playground/validator/v10"

type Username struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

func ensureIsValidUsername(username *Username) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(username)

	return
}

func NewUsername(username string) (*Username, error) {
	usernameVO := &Username{username}

	err := ensureIsValidUsername(usernameVO)

	if err != nil {
		return new(Username), err
	}

	return usernameVO, nil
}
