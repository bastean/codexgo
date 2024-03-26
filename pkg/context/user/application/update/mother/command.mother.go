package update

import (
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *update.Command {
	return update.NewCommand(valueObjectMother.RandomId().Value, valueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *update.Command {
	return update.NewCommand(valueObjectMother.InvalidId().Value, valueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
