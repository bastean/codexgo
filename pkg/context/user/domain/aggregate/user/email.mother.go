package user

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func EmailWithValidValue() *Email {
	value, err := values.New[*Email](services.Create.Email())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func EmailWithInvalidValue() (string, error) {
	spaces := strings.Repeat(" ", services.Create.IntRange(1, 12))

	value := services.Create.RandomString([]string{
		services.Create.LoremIpsumWord(),
		spaces + services.Create.Email(),
		services.Create.Email() + spaces,
		spaces + services.Create.Email() + spaces,
	})

	_, err := values.New[*Email](value)

	return value, err
}
