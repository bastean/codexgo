package create

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

func RandomCommand() *Command {
	id := user.IdWithValidValue()
	email := user.EmailWithValidValue()
	username := user.UsernameWithValidValue()
	password := user.PasswordWithValidValue()

	return &Command{
		Id:       id.Value,
		Email:    email.Value,
		Username: username.Value,
		Password: password.Value,
	}
}
