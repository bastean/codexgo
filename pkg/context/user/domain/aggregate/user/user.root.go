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
	Created, Updated              string
	ID, Email, Username, Password string
	Verified                      bool
}

func (u *User) ToPrimitive() *Primitive {
	return &Primitive{
		Created:  u.Created.Value,
		Updated:  u.Updated.Value,
		ID:       u.ID.Value,
		Email:    u.Email.Value,
		Username: u.Username.Value,
		Password: u.CipherPassword.Value,
		Verified: u.Verified.Value,
	}
}

func (u *User) IsVerified() bool {
	return u.Verified.Value
}

func create(user *Primitive) (*User, error) {
	root, errRoot := aggregates.NewRoot()

	id, errID := NewID(user.ID)
	email, errEmail := NewEmail(user.Email)
	username, errUsername := NewUsername(user.Username)
	verified, errVerified := NewVerified(user.Verified)

	if err := errors.Join(errRoot, errID, errEmail, errUsername, errVerified); err != nil {
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
	aggregate, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	created, errCreated := aggregates.NewTime(primitive.Created)
	updated, errUpdated := aggregates.NewTime(primitive.Updated)
	cipherPassword, errCipherPassword := NewCipherPassword(primitive.Password)

	if err := errors.Join(errCreated, errUpdated, errCipherPassword); err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

	aggregate.Created = created
	aggregate.Updated = updated
	aggregate.CipherPassword = cipherPassword

	return aggregate, nil
}

func FromRaw(raw *Primitive) (*User, error) {
	raw.Verified = false

	aggregate, err := create(raw)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromRaw")
	}

	aggregate.PlainPassword, err = NewPlainPassword(raw.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromRaw")
	}

	return aggregate, nil
}

func New(raw *Primitive) (*User, error) {
	aggregate, err := FromRaw(raw)

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
