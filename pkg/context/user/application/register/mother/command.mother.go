package commandMother

import (
	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *register.Command {
	return register.NewCommand(sharedValueObjectMother.RandomId().Value, sharedValueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *register.Command {
	return register.NewCommand(sharedValueObjectMother.InvalidId().Value, sharedValueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
