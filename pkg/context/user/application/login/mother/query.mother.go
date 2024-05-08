package queryMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *login.Query {
	email, _ := valueObjectMother.RandomEmail()
	password, _ := valueObjectMother.RandomPassword()

	return &login.Query{
		Email:    email.Value(),
		Password: password.Value(),
	}
}
