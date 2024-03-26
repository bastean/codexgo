package queryMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *login.Query {
	return login.NewQuery(valueObjectMother.RandomEmail().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *login.Query {
	return login.NewQuery(valueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
