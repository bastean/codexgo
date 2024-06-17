package update

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.IdWithValidValue()
	email := valueobj.EmailWithValidValue()
	username := valueobj.UsernameWithValidValue()
	password := valueobj.PasswordWithValidValue()
	updatedPassword := valueobj.PasswordWithValidValue()

	return &Command{
		Id:              id.Value(),
		Email:           email.Value(),
		Username:        username.Value(),
		Password:        password.Value(),
		UpdatedPassword: updatedPassword.Value(),
	}
}
