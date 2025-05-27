package values

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type Email struct {
	String
}

func (e *Email) Validate() error {
	if IsNotValid(e.RawValue(), "startsnotwith= ", "endsnotwith= ", "email") {
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
