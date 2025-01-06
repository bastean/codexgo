package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	QueryMinCharactersLength = "1"
	QueryMaxCharactersLength = "20"
)

type Query struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewQuery(value string) (*Query, error) {
	value = strings.TrimSpace(value)

	object := &Query{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewQuery",
			What:  fmt.Sprintf("Query must be between %s to %s characters and be alpha only", QueryMinCharactersLength, QueryMaxCharactersLength),
			Why: errors.Meta{
				"Query": value,
			},
		})
	}

	return object, nil
}
