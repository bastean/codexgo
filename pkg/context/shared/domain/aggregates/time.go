package aggregates

import (
	"time"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type Time struct {
	Value string
}

func NewTime(value string) (*Time, error) {
	if _, err := time.Parse(time.RFC3339Nano, value); err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "NewTime",
			What:  "Invalid Time format",
			Why: errors.Meta{
				"Time": value,
			},
		})
	}

	return &Time{
		Value: value,
	}, nil
}
