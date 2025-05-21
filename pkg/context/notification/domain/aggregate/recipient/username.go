package recipient

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

const (
	UsernameMinCharactersLength = "2"
	UsernameMaxCharactersLength = "20"
)

type Username struct {
	values.Object[string]
}

func (u *Username) Validate() error {
	if values.IsNotValid(u.RawValue(), "startsnotwith= ", "endsnotwith= ", "gte=2", "lte=20", "alphanum") {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: fmt.Sprintf("Username must be between %s to %s characters and be alphanumeric only", UsernameMinCharactersLength, UsernameMaxCharactersLength),
			Why: errors.Meta{
				"Username": u.RawValue(),
			},
		})
	}

	u.Valid()

	return nil
}
