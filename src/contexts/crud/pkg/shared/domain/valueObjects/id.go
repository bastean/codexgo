package valueObjects

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Id struct {
	Value string `validate:"uuid"`
}

func ensureIsValidId(id *Id) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(id)

	return
}

func NewId(id string) (*Id, error) {
	id = strings.TrimSpace(id)
	idVO := &Id{id}

	err := ensureIsValidId(idVO)

	if err != nil {
		return new(Id), err
	}

	return idVO, nil
}
