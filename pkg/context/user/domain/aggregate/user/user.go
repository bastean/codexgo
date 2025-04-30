package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
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
		Created:  u.Created.Value(),
		Updated:  u.Updated.Value(),
		ID:       u.ID.Value(),
		Email:    u.Email.Value(),
		Username: u.Username.Value(),
		Password: u.CipherPassword.Value(),
		Verified: u.Verified.Value(),
	}

	if u.Verify != nil {
		primitive.Verify = u.Verify.Value()
	}

	if u.Reset != nil {
		primitive.Reset = u.Reset.Value()
	}

	return primitive
}

func (u *User) IsVerified() bool {
	return u.Verified.Value()
}

func (u *User) HasReset() bool {
	return u.Reset != nil
}

func (u *User) ValidateVerify(token *ID) error {
	if u.Verify.Value() != token.Value() {
		return errors.New[errors.Failure](&errors.Bubble{
			What: "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value(),
			},
		})
	}

	return nil
}

func (u *User) ValidateReset(token *ID) error {
	if u.Reset.Value() != token.Value() {
		return errors.New[errors.Failure](&errors.Bubble{
			What: "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value(),
			},
		})
	}

	return nil
}

func create(user *Primitive) (*User, error) {
	root, errRoot := aggregates.NewRoot()

	id, errID := values.New[*ID](user.ID)
	email, errEmail := values.New[*Email](user.Email)
	username, errUsername := values.New[*Username](user.Username)
	verified, errVerified := values.New[*Verified](user.Verified)

	if err := errors.Join(errRoot, errID, errEmail, errUsername, errVerified); err != nil {
		return nil, errors.BubbleUp(err)
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
		return nil, errors.BubbleUp(err)
	}

	var errCreated, errUpdated, errCipherPassword, errVerify, errReset error

	aggregate.Created, errCreated = values.New[*aggregates.Time](primitive.Created)
	aggregate.Updated, errUpdated = values.New[*aggregates.Time](primitive.Updated)
	aggregate.CipherPassword, errCipherPassword = values.New[*CipherPassword](primitive.Password)

	if primitive.Verify != "" {
		aggregate.Verify, errVerify = values.New[*ID](primitive.Verify)
	}

	if primitive.Reset != "" {
		aggregate.Reset, errReset = values.New[*ID](primitive.Reset)
	}

	if err := errors.Join(errCreated, errUpdated, errCipherPassword, errVerify, errReset); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return aggregate, nil
}

func FromRaw(raw *Primitive) (*User, error) {
	raw.Verified = false

	aggregate, err := create(raw)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate.PlainPassword, err = values.New[*PlainPassword](raw.Password)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return aggregate, nil
}

func New(raw *Primitive) (*User, error) {
	aggregate, err := FromRaw(raw)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate.Verify, err = values.New[*ID](raw.Verify)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate.Record(messages.New(
		events.UserCreatedSucceededKey,
		&events.UserCreatedSucceededAttributes{
			Verify:   aggregate.Verify.Value(),
			ID:       aggregate.ID.Value(),
			Email:    aggregate.Email.Value(),
			Username: aggregate.Username.Value(),
		},
		new(events.UserCreatedSucceededMeta),
	))

	return aggregate, nil
}
