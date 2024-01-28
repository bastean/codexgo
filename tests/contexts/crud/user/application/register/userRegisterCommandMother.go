package register

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/user/application/register"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObject"
)

func Create(id *sharedVO.Id, email *sharedVO.Email, username *userVO.Username, password *userVO.Password) *register.Command {
	return &register.Command{Id: id.Value, Email: email.Value, Username: username.Value, Password: password.Value}
}

func Random() *register.Command {
	return Create(create.RandomId(), create.RandomEmail(), create.RandomUsername(), create.RandomPassword())
}

func Invalid() *register.Command {
	return Create(create.InvalidId(), create.InvalidEmail(), create.WithInvalidUsernameLength(), create.WithInvalidPasswordLength())
}
