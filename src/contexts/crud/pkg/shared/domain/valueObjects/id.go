package valueObjects

import (
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	Value string `validate:"uuid4"`
}

var InvalidIdValue = errors.InvalidValue{Message: "Id value is invalid"}

func ensureIsValidId(id *Id) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(id)

	return
}

func NewId(id string) *Id {
	id = strings.TrimSpace(id)
	idVO := &Id{id}

	err := ensureIsValidId(idVO)

	if err != nil {
		panic(InvalidIdValue)
	}

	return idVO
}
