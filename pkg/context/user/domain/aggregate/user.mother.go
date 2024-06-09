package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomUser() *User {
	id, _ := valueobj.IdWithValidValue()
	email, _ := valueobj.EmailWithValidValue()
	username, _ := valueobj.UsernameWithValidValue()
	password, _ := valueobj.PasswordWithValidValue()

	user, _ := NewUser(&UserPrimitive{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	})

	return user
}
