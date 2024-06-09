package update

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.IdWithValidValue()
	email, _ := valueobj.EmailWithValidValue()
	username, _ := valueobj.UsernameWithValidValue()
	password, _ := valueobj.PasswordWithValidValue()
	updatedPassword, _ := valueobj.PasswordWithValidValue()

	return &Command{
		Id:              id.Value(),
		Email:           email.Value(),
		Username:        username.Value(),
		Password:        password.Value(),
		UpdatedPassword: updatedPassword.Value(),
	}
}
