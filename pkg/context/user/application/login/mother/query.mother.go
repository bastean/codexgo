package queryMother

import (
	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *login.Query {
	return login.NewQuery(sharedValueObjectMother.RandomEmail().Value, valueObjectMother.RandomPassword().Value)
}

func Invalid() *login.Query {
	return login.NewQuery(sharedValueObjectMother.InvalidEmail().Value, valueObjectMother.WithInvalidPasswordLength().Value)
}
