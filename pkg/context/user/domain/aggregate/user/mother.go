package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) PlainPasswordValid() *PlainPassword {
	value, err := values.New[*PlainPassword](m.Regex(`^[\W\w]{8,64}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) PlainPasswordInvalidLength() (string, error) {
	value := "x"

	_, err := values.New[*PlainPassword](value)

	return value, err
}

func (m *m) PasswordReplace(old *Password, value string) *Password {
	password, err := values.Replace(old, value)

	if err != nil {
		errors.Panic(err)
	}

	return password
}

func (m *m) PasswordValid() *Password {
	value, err := values.New[*Password](m.Regex(`^.{8,64}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) PasswordInvalid() (string, error) {
	var value string

	_, err := values.New[*Password](value)

	return value, err
}

func (m *m) VerifiedReplace(old *Verified, value bool) *Verified {
	verified, err := values.Replace(old, value)

	if err != nil {
		errors.Panic(err)
	}

	return verified
}

func (m *m) VerifiedValid() *Verified {
	value, err := values.New[*Verified](m.Bool())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) VerifiedTrue() *Verified {
	value, err := values.New[*Verified](true)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) VerifiedFalse() *Verified {
	value, err := values.New[*Verified](false)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) UserCopy(user *User) *User {
	copy, err := FromPrimitive(user.ToPrimitive())

	if err != nil {
		errors.Panic(err)
	}

	return copy
}

func (m *m) UserValid() *User {
	user, err := New(&Required{
		VerifyToken: values.Mother().IDValid().Value(),
		ID:          values.Mother().IDValid().Value(),
		Email:       values.Mother().EmailValid().Value(),
		Username:    values.Mother().UsernameValid().Value(),
		Password:    m.PasswordValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

func (m *m) UserValidFromPrimitive(without ...string) *User {
	user, err := FromPrimitive(&Primitive{
		CreatedAt:   values.Mother().TimeValid().ToPrimitive(),
		UpdatedAt:   values.Mother().TimeValid().ToPrimitive(),
		VerifyToken: token.Mother().TokenValid().ToPrimitive(),
		ResetToken:  token.Mother().TokenValid().ToPrimitive(),
		ID:          values.Mother().IDValid().ToPrimitive(),
		Email:       values.Mother().EmailValid().ToPrimitive(),
		Username:    values.Mother().UsernameValid().ToPrimitive(),
		Password:    m.PasswordValid().ToPrimitive(),
		Verified:    m.VerifiedValid().ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err)
	}

	for _, field := range without {
		switch field {
		case "UpdatedAt":
			user.UpdatedAt = nil
		case "VerifyToken":
			user.VerifyToken = nil
		case "ResetToken":
			user.ResetToken = nil
		default:
			errors.Panic(errors.Standard("Unknown field %q", field))
		}
	}

	return user
}

var Mother = mother.New[m]
