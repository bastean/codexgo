package aggregateMother

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *aggregate.User {
	id, _ := valueObjectMother.RandomId()
	email, _ := valueObjectMother.RandomEmail()
	username, _ := valueObjectMother.RandomUsername()
	password, _ := valueObjectMother.RandomPassword()

	user, _ := aggregate.NewUser(id.Value(), email.Value(), username.Value(), password.Value())

	return user
}
