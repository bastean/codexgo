package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	value string `validate:"uuid4"`
}

func (id *Id) Value() string {
	return id.value
}

func (id *Id) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(id)
}

func NewId(id string) (model.ValueObject[string], error) {
	id = strings.TrimSpace(id)

	idVO := &Id{
		value: id,
	}

	if idVO.IsValid() != nil {
		return nil, errs.NewInvalidValueError(&errs.Bubble{
			Where: "NewId",
			What:  "invalid format",
			Why: errs.Meta{
				"Id": id,
			},
		})
	}

	return idVO, nil
}
