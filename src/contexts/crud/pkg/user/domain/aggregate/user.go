package aggregate

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type User struct {
	*sharedVO.Id
	*sharedVO.Email
	*userVO.Username
	*userVO.Password
}

type UserPrimitive struct {
	Id       string
	Email    string
	Username string
	Password string
}

func newUser(id, email, username, password string) (*User, error) {

	idVO := sharedVO.NewId(id)
	emailVO := sharedVO.NewEmail(email)
	usernameVO := userVO.NewUsername(username)
	passwordVO := userVO.NewPassword(password)

	return &User{
		idVO,
		emailVO,
		usernameVO,
		passwordVO,
	}, nil
}

func (user *User) ToPrimitives() *UserPrimitive {
	return &UserPrimitive{
		user.Id.Value,
		user.Email.Value,
		user.Username.Value,
		user.Password.Value,
	}
}

func FromPrimitives(userPrimitive *UserPrimitive) (user *User, err error) {
	user, err = newUser(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
	)

	return
}

func Create(id, email, username, password string) (user *User, err error) {
	user, err = newUser(
		id,
		email,
		username,
		password,
	)

	return
}
