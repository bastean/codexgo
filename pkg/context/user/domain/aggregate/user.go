package aggregate

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/message"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type User struct {
	*aggregate.AggregateRoot
	Id       model.ValueObject[string]
	Email    model.ValueObject[string]
	Username model.ValueObject[string]
	Password model.ValueObject[string]
	Verified model.ValueObject[bool]
}

type UserPrimitive struct {
	Id       string
	Email    string
	Username string
	Password string
	Verified bool
}

func create(id, email, username, password string, verified bool) (*User, error) {
	aggregateRoot := aggregate.NewAggregateRoot()

	idVO, idErr := valueObject.NewId(id)
	emailVO, emailErr := valueObject.NewEmail(email)
	usernameVO, usernameErr := valueObject.NewUsername(username)
	passwordVO, passwordErr := valueObject.NewPassword(password)
	verifiedVO, verifiedErr := valueObject.NewVerified(verified)

	err := errors.Join(idErr, emailErr, usernameErr, passwordErr, verifiedErr)

	if err != nil {
		return nil, errs.BubbleUp("create", err)
	}

	return &User{
		AggregateRoot: aggregateRoot,
		Id:            idVO,
		Email:         emailVO,
		Username:      usernameVO,
		Password:      passwordVO,
		Verified:      verifiedVO,
	}, nil
}

func (user *User) ToPrimitives() *UserPrimitive {
	return &UserPrimitive{
		Id:       user.Id.Value(),
		Email:    user.Email.Value(),
		Username: user.Username.Value(),
		Password: user.Password.Value(),
		Verified: user.Verified.Value(),
	}
}

func FromPrimitives(userPrimitive *UserPrimitive) (*User, error) {
	user, err := create(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
		userPrimitive.Verified,
	)

	if err != nil {
		return nil, errs.BubbleUp("FromPrimitives", err)
	}

	return user, nil
}

func NewUser(id, email, username, password string) (*User, error) {
	verified := false

	user, err := create(
		id,
		email,
		username,
		password,
		verified,
	)

	if err != nil {
		return nil, errs.BubbleUp("NewUser", err)
	}

	eventMessage, err := message.NewRegisteredSucceededEvent(&message.RegisteredSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	})

	if err != nil {
		return nil, errs.BubbleUp("NewUser", err)
	}

	user.RecordMessage(eventMessage)

	return user, nil
}
