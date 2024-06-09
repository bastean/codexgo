package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
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

func NewId(id string) (models.ValueObject[string], error) {
	id = strings.TrimSpace(id)

	idVO := &Id{
		Id: id,
	}

	if idVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewId",
			What:  "invalid uuid4 format",
			Why: errors.Meta{
				"Id": id,
			},
		})
	}

	return idVO, nil
}
