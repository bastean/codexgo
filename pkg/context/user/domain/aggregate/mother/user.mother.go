package aggregateMother

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func FromCommand(command *register.Command) *aggregate.User {
	return aggregate.NewUser(sharedValueObject.NewId(command.Id).Value, sharedValueObject.NewEmail(command.Email).Value, valueObject.NewUsername(command.Username).Value, valueObject.NewPassword(command.Password).Value)
}

func Invalid() *aggregate.User {
	return aggregate.NewUser(sharedValueObjectMother.InvalidId().Value, sharedValueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}

func Random() *aggregate.User {
	return aggregate.NewUser(sharedValueObjectMother.RandomId().Value, sharedValueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value)
}
