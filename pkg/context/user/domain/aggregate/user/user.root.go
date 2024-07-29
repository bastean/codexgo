package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type User struct {
	*aggregates.Root
	*Id
	*Email
	*Username
	*Password
	*Verified
}

type Primitive struct {
	Id, Email, Username, Password string
	Verified                      bool
}

func create(primitive *Primitive) (*User, error) {
	root := aggregates.NewRoot()

	id, errId := NewId(primitive.Id)
	email, errEmail := NewEmail(primitive.Email)
	username, errUsername := NewUsername(primitive.Username)
	password, errPassword := NewPassword(primitive.Password)
	verified, errVerified := NewVerified(primitive.Verified)

	if err := errors.Join(errId, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Root:     root,
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
		Verified: verified,
	}, nil
}

func (user *User) ToPrimitive() *Primitive {
	return &Primitive{
		Id:       user.Id.Value,
		Email:    user.Email.Value,
		Username: user.Username.Value,
		Password: user.Password.Value,
		Verified: user.Verified.Value,
	}
}

func FromPrimitive(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	return user, nil
}

func New(primitive *Primitive) (*User, error) {
	primitive.Verified = false

	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	message, err := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			Id:       user.Id.Value,
			Email:    user.Email.Value,
			Username: user.Username.Value,
		},
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	user.RecordMessage(message)

	return user, nil
}
