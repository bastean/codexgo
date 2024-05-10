package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomUser() *User {
	id, _ := valueobj.RandomId()
	email, _ := valueobj.RandomEmail()
	username, _ := valueobj.RandomUsername()
	password, _ := valueobj.RandomPassword()

	user, _ := NewUser(id.Value(), email.Value(), username.Value(), password.Value())

	return user
}
