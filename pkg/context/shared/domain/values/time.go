package values

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type Time struct {
	String
}

func (t *Time) Validate() error {
	_ = time.Parse(t.RawValue())

	t.Valid()

	return nil
}
