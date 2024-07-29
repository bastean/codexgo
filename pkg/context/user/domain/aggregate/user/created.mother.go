package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

func RandomCreatedSucceeded() *messages.Message {
	id := IdWithValidValue()
	email := EmailWithValidValue()
	username := UsernameWithValidValue()

	event, err := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			Id:       id.Value,
			Email:    email.Value,
			Username: username.Value,
		},
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomCreatedSucceeded")
	}

	return event
}
