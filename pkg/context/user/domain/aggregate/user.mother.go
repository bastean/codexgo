package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomUser() *User {
	id, _ := valueobj.RandomId()
	email, _ := valueobj.RandomEmail()
	username, _ := valueobj.RandomUsername()
	password, _ := valueobj.RandomPassword()

	user, _ := NewUser(&UserPrimitive{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	})

	return user
}
