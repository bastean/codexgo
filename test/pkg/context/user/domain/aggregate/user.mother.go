package aggregate

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	user "github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	create "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
)

func Create(id *sharedVO.Id, email *sharedVO.Email, username *userVO.Username, password *userVO.Password) *user.User {
	return user.Create(id.Value, email.Value, username.Value, password.Value)
}

func FromCommand(command register.Command) *user.User {
	return Create(create.NewId(command.Id), create.NewEmail(command.Email), create.NewUsername(command.Username), create.NewPassword(command.Password))
}

func Invalid() *user.User {
	return Create(create.InvalidId(), create.InvalidEmail(), create.WithInvalidUsernameLength(), create.WithInvalidPasswordLength())
}

func Random() *user.User {
	return Create(create.RandomId(), create.RandomEmail(), create.RandomUsername(), create.RandomPassword())
}
