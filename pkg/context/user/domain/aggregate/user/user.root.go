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
	*PlainPassword
	*CipherPassword
	*Verified
}

type Primitive struct {
	ID, Email, Username, Password string
	Verified                      bool
}

func (u *User) ToPrimitive() *Primitive {
	return &Primitive{
		ID:       u.ID.Value,
		Email:    u.Email.Value,
		Username: u.Username.Value,
		Password: u.CipherPassword.Value,
		Verified: u.Verified.Value,
	}
}

func create(user *Primitive) (*User, error) {
	root := aggregates.NewRoot()

	id, errID := NewID(user.ID)
	email, errEmail := NewEmail(user.Email)
	username, errUsername := NewUsername(user.Username)
	verified, errVerified := NewVerified(user.Verified)

	if err := errors.Join(errID, errEmail, errUsername, errVerified); err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Root:     root,
		ID:       id,
		Email:    email,
		Username: username,
		Verified: verified,
	}, nil
}

func FromPrimitive(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	password, err := NewCipherPassword(primitive.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	user.CipherPassword = password

	return user, nil
}

func FromRaw(raw *Primitive) (*User, error) {
	raw.Verified = false

	aggregate, err := create(raw)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromRaw")
	}

	password, err := NewPlainPassword(raw.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromRaw")
	}

	aggregate.PlainPassword = password

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
