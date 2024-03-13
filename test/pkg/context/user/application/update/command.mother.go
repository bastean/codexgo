package update

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	valueObjectMother "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
)

func Create(id *sharedVO.Id, email *sharedVO.Email, username *userVO.Username, password *userVO.Password, updatedPassword *userVO.Password) *update.Command {
	return &update.Command{Id: id.Value, Email: email.Value, Username: username.Value, Password: password.Value, UpdatedPassword: updatedPassword.Value}
}

func Random() *update.Command {
	return Create(valueObjectMother.RandomId(), valueObjectMother.RandomEmail(), valueObjectMother.RandomUsername(), valueObjectMother.RandomPassword(), valueObjectMother.RandomPassword())
}

func Invalid() *update.Command {
	return Create(valueObjectMother.InvalidId(), valueObjectMother.InvalidEmail(), valueObjectMother.WithInvalidUsernameLength(), valueObjectMother.WithInvalidPasswordLength(), valueObjectMother.WithInvalidPasswordLength())
}
