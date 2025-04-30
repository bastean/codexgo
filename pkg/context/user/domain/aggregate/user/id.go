package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type ID struct {
	values.Object[string]
}

func (id *ID) Validate() error {
	if services.IsNotValid(id.RawValue(), "startsnotwith= ", "endsnotwith= ", "uuid4") {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Invalid UUID4 format",
			Why: errors.Meta{
				"ID": id.RawValue(),
			},
		})
	}

	id.Valid()

	return nil
}
