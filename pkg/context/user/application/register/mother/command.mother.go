package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *register.Command {
	id, _ := valueObjectMother.RandomId()
	email, _ := valueObjectMother.RandomEmail()
	username, _ := valueObjectMother.RandomUsername()
	password, _ := valueObjectMother.RandomPassword()

	return &register.Command{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	}
}
