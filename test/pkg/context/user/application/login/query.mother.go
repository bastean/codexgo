package login

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	create "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
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
