package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type User struct {
	*aggregates.Root
	VerifyToken, ResetToken *ID
	*ID
	*Email
	*Username
	*Password
	*Verified
}

type Primitive struct {
	Created, Updated        *values.Primitive[string]
	VerifyToken, ResetToken *values.Primitive[string]
	ID, Email, Username     *values.Primitive[string]
	Password                *values.Primitive[string]
	Verified                *values.Primitive[bool]
}

type Criteria struct {
	*ID
	*Email
	*Username
}

type Required struct {
	VerifyToken         string
	ID, Email, Username string
	Password            string
}

func (u *User) IsVerified() bool {
	return u.Verified.Value()
}

func (u *User) HasResetToken() bool {
	return u.ResetToken != nil
}

func (u *User) ValidateVerifyToken(token *ID) error {
	if u.VerifyToken.Value() != token.Value() {
		return errors.New[errors.Failure](&errors.Bubble{
			What: "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value(),
			},
		})
	}

	return nil
}

func (u *User) ValidateResetToken(token *ID) error {
	if u.ResetToken.Value() != token.Value() {
		return errors.New[errors.Failure](&errors.Bubble{
			What: "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value(),
			},
		})
	}

	return nil
}

func (u *User) ToPrimitive() *Primitive {
	primitive := &Primitive{
		Created:  u.Created.ToPrimitive(),
		Updated:  u.Updated.ToPrimitive(),
		ID:       u.ID.ToPrimitive(),
		Email:    u.Email.ToPrimitive(),
		Username: u.Username.ToPrimitive(),
		Password: u.Password.ToPrimitive(),
		Verified: u.Verified.ToPrimitive(),
	}

	if u.VerifyToken != nil {
		primitive.VerifyToken = u.VerifyToken.ToPrimitive()
	}

	if u.ResetToken != nil {
		primitive.ResetToken = u.ResetToken.ToPrimitive()
	}

	return primitive
}

func FromPrimitive(primitive *Primitive) (*User, error) {
	created, errCreated := values.FromPrimitive[*aggregates.Time](primitive.Created)
	updated, errUpdated := values.FromPrimitive[*aggregates.Time](primitive.Updated, true)

	verifyToken, errVerifyToken := values.FromPrimitive[*ID](primitive.VerifyToken, true)
	resetToken, errResetToken := values.FromPrimitive[*ID](primitive.ResetToken, true)

	id, errID := values.FromPrimitive[*ID](primitive.ID)
	email, errEmail := values.FromPrimitive[*Email](primitive.Email)
	username, errUsername := values.FromPrimitive[*Username](primitive.Username)
	password, errPassword := values.FromPrimitive[*Password](primitive.Password)
	verified, errVerified := values.FromPrimitive[*Verified](primitive.Verified)

	if err := errors.Join(errCreated, errUpdated, errVerifyToken, errResetToken, errID, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return &User{
		Root: &aggregates.Root{
			Created: created,
			Updated: updated,
			Events:  make([]*messages.Message, 0),
		},
		VerifyToken: verifyToken,
		ResetToken:  resetToken,
		ID:          id,
		Email:       email,
		Username:    username,
		Password:    password,
		Verified:    verified,
	}, nil
}

func New(required *Required) (*User, error) {
	verifyToken, errVerifyToken := values.New[*ID](required.VerifyToken)
	id, errID := values.New[*ID](required.ID)

	email, errEmail := values.New[*Email](required.Email)
	username, errUsername := values.New[*Username](required.Username)
	password, errPassword := values.New[*Password](required.Password)
	verified, errVerified := values.New[*Verified](false)

	if err := errors.Join(errVerifyToken, errID, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err)
	}

	user := &User{
		Root: &aggregates.Root{
			Events: make([]*messages.Message, 0),
		},
		VerifyToken: verifyToken,
		ID:          id,
		Email:       email,
		Username:    username,
		Password:    password,
		Verified:    verified,
	}

	user.Record(messages.New(
		CreatedSucceededKey,
		&CreatedSucceededAttributes{
			VerifyToken: user.VerifyToken.Value(),
			ID:          user.ID.Value(),
			Email:       user.Email.Value(),
			Username:    user.Username.Value(),
		},
		new(CreatedSucceededMeta),
	))

	return user, nil
}
