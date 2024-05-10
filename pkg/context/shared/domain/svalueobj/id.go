package svalueobj

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	Id string `validate:"uuid4"`
}

func (id *Id) Value() string {
	return id.Id
}

func (id *Id) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(id)
}

func NewId(id string) (smodel.ValueObject[string], error) {
	id = strings.TrimSpace(id)

	idVO := &Id{
		Id: id,
	}

	if idVO.IsValid() != nil {
		return nil, serror.NewInvalidValueError(&serror.Bubble{
			Where: "NewId",
			What:  "invalid format",
			Why: serror.Meta{
				"Id": id,
			},
		})
	}

	return idVO, nil
}
