package errs

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/codexgo/internal/pkg/service/errors"
)

func BindingJSON(who error, where string) error {
	var err *json.UnmarshalTypeError

	if errors.As(who, &err) {
		return errors.NewFailure(
			where,
			fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", err.Field, err.Type, err.Value),
			who,
		)
	}

	return errors.NewInternal(
		where,
		"Cannot bind a JSON to a struct",
		who,
	)
}
