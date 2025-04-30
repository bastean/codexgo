package aggregates

import (
	"time"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Time struct {
	values.Object[string]
}

func (t *Time) Validate() error {
	if _, err := time.Parse(time.RFC3339Nano, t.RawValue()); err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Invalid Time format",
			Why: errors.Meta{
				"Time": t.RawValue(),
			},
		})
	}

	t.Valid()

	return nil
}
