package aggregateMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func FromCommand(command *register.Command) *aggregate.User {
	return aggregate.NewUser(valueObject.NewId(command.Id).Value, valueObject.NewEmail(command.Email).Value, valueObject.NewUsername(command.Username).Value, valueObject.NewPassword(command.Password).Value)
}

func Invalid() *aggregate.User {
	return aggregate.NewUser(valueObjectMother.InvalidId().Value, valueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}

func Random() *aggregate.User {
	return aggregate.NewUser(valueObjectMother.RandomId().Value, valueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value)
}
