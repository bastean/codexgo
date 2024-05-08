package update

import (
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *update.Command {
	id, _ := valueObjectMother.RandomId()
	email, _ := valueObjectMother.RandomEmail()
	username, _ := valueObjectMother.RandomUsername()
	password, _ := valueObjectMother.RandomPassword()
	updatedPassword, _ := valueObjectMother.RandomPassword()

	return &update.Command{
		Id:              id.Value(),
		Email:           email.Value(),
		Username:        username.Value(),
		Password:        password.Value(),
		UpdatedPassword: updatedPassword.Value(),
	}
}
