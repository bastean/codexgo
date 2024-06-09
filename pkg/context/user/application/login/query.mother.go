package login

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	email, _ := valueobj.EmailWithValidValue()
	password, _ := valueobj.PasswordWithValidValue()

	return &Query{
		Email:    email.Value(),
		Password: password.Value(),
	}
}
