package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomUser() *User {
	id := valueobj.IdWithValidValue()
	email := valueobj.EmailWithValidValue()
	username := valueobj.UsernameWithValidValue()
	password := valueobj.PasswordWithValidValue()

	user, err := NewUser(&UserPrimitive{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomUser")
	}

	return user
}
