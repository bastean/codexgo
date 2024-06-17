package event

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCreatedSucceeded() *messages.Message {
	id := valueobj.IdWithValidValue()
	email := valueobj.EmailWithValidValue()
	username := valueobj.UsernameWithValidValue()

	event, err := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			Id:       id.Value(),
			Email:    email.Value(),
			Username: username.Value(),
		},
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomCreatedSucceeded")
	}

	return event
}
