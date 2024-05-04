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
	*valueObject.Verified
}

type UserPrimitive struct {
	Id       string
	Email    string
	Username string
	Password string
	Verified bool
}

func create(id, email, username, password string, verified bool) *User {
	aggregateRoot := aggregate.NewAggregateRoot()

	idVO := valueObject.NewId(id)
	emailVO := valueObject.NewEmail(email)
	usernameVO := valueObject.NewUsername(username)
	passwordVO := valueObject.NewPassword(password)
	verifiedVO := valueObject.NewVerified(verified)

	return &User{
		AggregateRoot: aggregateRoot,
		Id:            idVO,
		Email:         emailVO,
		Username:      usernameVO,
		Password:      passwordVO,
		Verified:      verifiedVO,
	}
}

func (user *User) ToPrimitives() *UserPrimitive {
	return &UserPrimitive{
		Id:       user.Id.Value,
		Email:    user.Email.Value,
		Username: user.Username.Value,
		Password: user.Password.Value,
		Verified: user.Verified.Value,
	}
}

func FromPrimitives(userPrimitive *UserPrimitive) *User {
	return create(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
		userPrimitive.Verified,
	)
}

func NewUser(id, email, username, password string) *User {
	verified := false

	user := create(
		id,
		email,
		username,
		password,
		verified,
	)

	user.RecordMessage(message.NewRegisteredSucceededEvent(&message.RegisteredSucceededEventAttributes{Id: user.Id.Value, Email: user.Email.Value, Username: user.Username.Value}))

	return user
}
