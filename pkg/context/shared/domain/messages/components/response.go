package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	ResponseMinCharactersLength = "1"
	ResponseMaxCharactersLength = "20"
)

type Response struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewResponse(value string) (*Response, error) {
	value = strings.TrimSpace(value)

	object := &Response{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewResponse",
			What:  fmt.Sprintf("Response must be between %s to %s characters and be alpha only", ResponseMinCharactersLength, ResponseMaxCharactersLength),
			Why: errors.Meta{
				"Response": value,
			},
		})
	}

	return object, nil
}
