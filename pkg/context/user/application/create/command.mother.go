package create

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()
	email, _ := valueobj.RandomEmail()
	username, _ := valueobj.RandomUsername()
	password, _ := valueobj.RandomPassword()

	return &Command{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	}
}
