package update

import (
	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *update.Command {
	return update.NewCommand(sharedValueObjectMother.RandomId().Value, sharedValueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *update.Command {
	return update.NewCommand(sharedValueObjectMother.InvalidId().Value, sharedValueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
