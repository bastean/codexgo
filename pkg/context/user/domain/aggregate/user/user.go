package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type User struct {
	*aggregates.Root
	Verify, Reset *ID
	*ID
	*Email
	*Username
	*PlainPassword
	*CipherPassword
	*Verified
}

type Primitive struct {
	Created, Updated              string
	Verify, Reset                 string
	ID, Email, Username, Password string
	Verified                      bool
}

type Criteria struct {
	*ID
	*Email
	*Username
}

func (u *User) ToPrimitive() *Primitive {
	primitive := &Primitive{
		Created:  u.Created.Value,
		Updated:  u.Updated.Value,
		ID:       u.ID.Value,
		Email:    u.Email.Value,
		Username: u.Username.Value,
		Password: u.CipherPassword.Value,
		Verified: u.Verified.Value,
	}

	if u.Verify != nil {
		primitive.Verify = u.Verify.Value
	}

	if u.Reset != nil {
		primitive.Reset = u.Reset.Value
	}

	return primitive
}

func (u *User) IsVerified() bool {
	return u.Verified.Value
}

func (u *User) HasReset() bool {
	return u.Reset != nil
}

func (u *User) ValidateVerify(token *ID) error {
	if u.Verify.Value != token.Value {
		return errors.New[errors.Failure](&errors.Bubble{
			Where: "ValidateVerify",
			What:  "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value,
			},
		})
	}

	return nil
}

func (u *User) ValidateReset(token *ID) error {
	if u.Reset.Value != token.Value {
		return errors.New[errors.Failure](&errors.Bubble{
			Where: "ValidateReset",
			What:  "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value,
			},
		})
	}

	return nil
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

	var errCreated, errUpdated, errCipherPassword, errVerify, errReset error

	aggregate.Created, errCreated = aggregates.NewTime(primitive.Created)
	aggregate.Updated, errUpdated = aggregates.NewTime(primitive.Updated)
	aggregate.CipherPassword, errCipherPassword = NewCipherPassword(primitive.Password)

	if primitive.Verify != "" {
		aggregate.Verify, errVerify = NewID(primitive.Verify)
	}

	if primitive.Reset != "" {
		aggregate.Reset, errReset = NewID(primitive.Reset)
	}

	if err := errors.Join(errCreated, errUpdated, errCipherPassword, errVerify, errReset); err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitive")
	}

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

	aggregate.Verify, err = NewID(raw.Verify)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	aggregate.Record(messages.New(
		events.UserCreatedSucceededKey,
		&events.UserCreatedSucceededAttributes{
			Verify:   aggregate.Verify.Value,
			ID:       aggregate.ID.Value,
			Email:    aggregate.Email.Value,
			Username: aggregate.Username.Value,
		},
		new(events.UserCreatedSucceededMeta),
	))

	return aggregate, nil
}
