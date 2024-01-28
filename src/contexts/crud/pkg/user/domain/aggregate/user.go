package aggregate

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
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

func newUser(id, email, username, password string) *User {
	idVO := sharedVO.NewId(id)
	emailVO := sharedVO.NewEmail(email)
	usernameVO := userVO.NewUsername(username)
	passwordVO := userVO.NewPassword(password)

	return &User{
		idVO,
		emailVO,
		usernameVO,
		passwordVO,
	}
}

func (user *User) ToPrimitives() *UserPrimitive {
	return &UserPrimitive{
		user.Id.Value,
		user.Email.Value,
		user.Username.Value,
		user.Password.Value,
	}
}

func FromPrimitives(userPrimitive *UserPrimitive) *User {
	return newUser(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
	)
}

func Create(id, email, username, password string) *User {
	return newUser(
		id,
		email,
		username,
		password,
	)
}
