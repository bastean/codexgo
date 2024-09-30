package errs

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func BindingJSON(who error, where string) error {
	var err *json.UnmarshalTypeError

	if errors.As(who, &err) {
		return errors.NewFailure(&errors.Bubble{
			Where: where,
			What:  fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", err.Field, err.Type, err.Value),
			Who:   who,
		})
	}

	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "Cannot bind a JSON to a struct",
		Who:   who,
	})
}
