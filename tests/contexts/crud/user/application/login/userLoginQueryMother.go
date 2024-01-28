package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObject"
)

func Create(email *sharedVO.Email, password *userVO.Password) *login.Query {
	return &login.Query{Email: email.Value, Password: password.Value}
}

func Random() *login.Query {
	return Create(create.RandomEmail(), create.RandomPassword())
}

func Invalid() *login.Query {
	return Create(create.InvalidEmail(), create.WithInvalidPasswordLength())
}
