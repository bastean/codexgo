package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *register.Command {
	return register.NewCommand(valueObjectMother.RandomId().Value, valueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *register.Command {
	return register.NewCommand(valueObjectMother.InvalidId().Value, valueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
