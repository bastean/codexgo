package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/message"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type User struct {
	*aggregate.AggregateRoot
	*valueObject.Id
	*valueObject.Email
	*valueObject.Username
	*valueObject.Password
}

type UserPrimitive struct {
	Id       string
	Email    string
	Username string
	Password string
}

func create(id, email, username, password string) *User {
	aggregateRoot := aggregate.NewAggregateRoot()

	idVO := valueObject.NewId(id)
	emailVO := valueObject.NewEmail(email)
	usernameVO := valueObject.NewUsername(username)
	passwordVO := valueObject.NewPassword(password)

	return &User{
		aggregateRoot,
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
	return create(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
	)
}

func NewUser(id, email, username, password string) *User {
	user := create(
		id,
		email,
		username,
		password,
	)

	user.RecordMessage(message.NewRegisteredSucceededEvent(&message.RegisteredSucceededEventAttributes{Id: user.Id.Value, Email: user.Email.Value, Username: user.Username.Value}))

	return user
}
