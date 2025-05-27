package values

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type ID struct {
	String
}

func (id *ID) Validate() error {
	if IsNotValid(id.RawValue(), "startsnotwith= ", "endsnotwith= ", "uuid4") {
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
