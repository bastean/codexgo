package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/root"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type User struct {
	*root.Root
	VerifyToken, ResetToken *token.Token
	*values.ID
	*values.Email
	*values.Username
	*Password
	*Verified
}

type Primitive struct {
	Created, Updated        *values.StringPrimitive
	VerifyToken, ResetToken *token.Primitive
	ID, Email, Username     *values.StringPrimitive
	Password                *values.StringPrimitive
	Verified                *values.BoolPrimitive
}

type Criteria struct {
	*values.ID
	*values.Email
	*values.Username
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

func (u *User) ValidateVerifyToken(token *token.Token) error {
	if u.VerifyToken.ID.Value() != token.ID.Value() {
		return errors.New[errors.Failure](&errors.Bubble{
			What: "Tokens do not match",
			Why: errors.Meta{
				"Received": token.Value(),
			},
		})
	}

	return nil
}

func (u *User) ValidateResetToken(token *token.Token) error {
	if u.ResetToken.ID.Value() != token.ID.Value() {
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
		ID:       u.ID.ToPrimitive(),
		Email:    u.Email.ToPrimitive(),
		Username: u.Username.ToPrimitive(),
		Password: u.Password.ToPrimitive(),
		Verified: u.Verified.ToPrimitive(),
	}

	if u.Updated != nil {
		primitive.Updated = u.Updated.ToPrimitive()
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
	created, errCreated := values.FromPrimitive[*values.Time](primitive.Created)
	updated, errUpdated := values.FromPrimitive[*values.Time](primitive.Updated, true)

	verifyToken, errVerifyToken := token.FromPrimitive(primitive.VerifyToken, true)
	resetToken, errResetToken := token.FromPrimitive(primitive.ResetToken, true)

	id, errID := values.FromPrimitive[*values.ID](primitive.ID)
	email, errEmail := values.FromPrimitive[*values.Email](primitive.Email)
	username, errUsername := values.FromPrimitive[*values.Username](primitive.Username)
	password, errPassword := values.FromPrimitive[*Password](primitive.Password)
	verified, errVerified := values.FromPrimitive[*Verified](primitive.Verified)

	if err := errors.Join(errCreated, errUpdated, errVerifyToken, errResetToken, errID, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return &User{
		Root: &root.Root{
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
	verifyToken, errVerifyToken := token.New(required.VerifyToken)
	id, errID := values.New[*values.ID](required.ID)

	email, errEmail := values.New[*values.Email](required.Email)
	username, errUsername := values.New[*values.Username](required.Username)
	password, errPassword := values.New[*Password](required.Password)
	verified, errVerified := values.New[*Verified](false)

	if err := errors.Join(errVerifyToken, errID, errEmail, errUsername, errPassword, errVerified); err != nil {
		return nil, errors.BubbleUp(err)
	}

	user := &User{
		Root: &root.Root{
			Events: make([]*messages.Message, 0),
		},
		VerifyToken: verifyToken,
		ID:          id,
		Email:       email,
		Username:    username,
		Password:    password,
		Verified:    verified,
	}

	err := user.CreationStamp()

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	user.Record(messages.New(
		CreatedSucceededKey,
		&CreatedSucceededAttributes{
			VerifyToken: user.VerifyToken.ID.Value(),
			ID:          user.ID.Value(),
			Email:       user.Email.Value(),
			Username:    user.Username.Value(),
		},
		new(CreatedSucceededMeta),
	))

	return user, nil
}
