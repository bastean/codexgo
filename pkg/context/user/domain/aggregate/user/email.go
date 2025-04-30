package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Email struct {
	values.Object[string]
}

func (e *Email) Validate() error {
	if services.IsNotValid(e.RawValue(), "startsnotwith= ", "endsnotwith= ", "email") {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Invalid email format",
			Why: errors.Meta{
				"Email": e.RawValue(),
			},
		})
	}

	e.Valid()

	return nil
}
