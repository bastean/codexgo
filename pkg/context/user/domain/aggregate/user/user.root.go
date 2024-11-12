package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type User struct {
	*aggregates.Root
	*ID
	*Email
	*Username
	*Password
	*Verified
}

type Primitive struct {
	ID, Email, Username, Password string
	Verified                      bool
}

func create(primitive *Primitive) (*User, error) {
	root := aggregates.NewRoot()

	id, errID := NewID(primitive.ID)
	email, errEmail := NewEmail(primitive.Email)
	username, errUsername := NewUsername(primitive.Username)
	password, errPassword := NewPassword(primitive.Password)
	verified, errVerified := NewVerified(primitive.Verified)

	if err := errors.Join(errID, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Root:     root,
		ID:       id,
		Email:    email,
		Username: username,
		Password: password,
		Verified: verified,
	}, nil
}

func (u *User) ToPrimitive() *Primitive {
	return &Primitive{
		ID:       u.ID.Value,
		Email:    u.Email.Value,
		Username: u.Username.Value,
		Password: u.Password.Value,
		Verified: u.Verified.Value,
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

	aggregate, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	aggregate.Record(messages.New[events.Event](
		user.CreatedSucceededKey,
		&user.CreatedSucceededAttributes{
			ID:       aggregate.ID.Value,
			Email:    aggregate.Email.Value,
			Username: aggregate.Username.Value,
		},
		&user.CreatedSucceededMeta{},
	))

	return aggregate, nil
}
