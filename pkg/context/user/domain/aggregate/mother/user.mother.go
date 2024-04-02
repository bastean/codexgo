package aggregateMother

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Invalid() *aggregate.User {
	return aggregate.NewUser(valueObjectMother.InvalidId().Value, valueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidUsernameLength().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}

func Random() *aggregate.User {
	return aggregate.NewUser(valueObjectMother.RandomId().Value, valueObjectMother.RandomEmail().Value, valueObjectMother.RandomUsername().Value, valueObjectMother.RandomPassword().Value)
}
