package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Time struct {
	values.Object[string]
}

func (t *Time) Validate() error {
	_ = time.Parse(t.RawValue())

	t.Valid()

	return nil
}
