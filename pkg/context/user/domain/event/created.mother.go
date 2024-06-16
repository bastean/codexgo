package event

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCreatedSucceeded() *messages.Message {
	id, _ := valueobj.IdWithValidValue()
	email, _ := valueobj.EmailWithValidValue()
	username, _ := valueobj.UsernameWithValidValue()

	event, _ := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			Id:       id.Value(),
			Email:    email.Value(),
			Username: username.Value(),
		},
	})

	return event
}
