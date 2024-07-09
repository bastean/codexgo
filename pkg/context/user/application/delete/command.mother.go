package delete

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

func RandomCommand() *Command {
	id := user.IdWithValidValue()
	password := user.PasswordWithValidValue()

	return &Command{
		Id:       id.Value,
		Password: password.Value,
	}
}
