package user

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) IDNew(id string) *ID {
	value, err := values.New[*ID](id)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IDValid() *ID {
	value, err := values.New[*ID](m.UUID())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IDInvalid() (string, error) {
	value := "x"

	_, err := values.New[*ID](value)

	return value, err
}

func (m *m) EmailNew(email string) *Email {
	value, err := values.New[*Email](email)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) EmailValid() *Email {
	value, err := values.New[*Email](m.Email())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) EmailInvalid() (string, error) {
	spaces := strings.Repeat(" ", m.IntRange(1, 12))

	value := m.RandomString([]string{
		m.LoremIpsumWord(),
		spaces + m.Email(),
		m.Email() + spaces,
		spaces + m.Email() + spaces,
	})

	_, err := values.New[*Email](value)

	return value, err
}

func (m *m) UsernameNew(username string) *Username {
	value, err := values.New[*Username](username)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) UsernameValid() *Username {
	value, err := values.New[*Username](m.Regex(`^[A-Za-z0-9]{2,20}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) UsernameInvalidLength() (string, error) {
	value := "x"

	_, err := values.New[*Username](value)

	return value, err
}

func (m *m) UsernameInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := values.New[*Username](value)

	return value, err
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

func (m *m) UserValidFromPrimitive() *User {
	user, err := FromPrimitive(&Primitive{
		Created:     aggregates.Mother().TimeValid().ToPrimitive(),
		Updated:     aggregates.Mother().TimeValid().ToPrimitive(),
		VerifyToken: m.IDValid().ToPrimitive(),
		ResetToken:  m.IDValid().ToPrimitive(),
		ID:          m.IDValid().ToPrimitive(),
		Email:       m.EmailValid().ToPrimitive(),
		Username:    m.UsernameValid().ToPrimitive(),
		Password:    m.PasswordValid().ToPrimitive(),
		Verified:    m.VerifiedValid().ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

func (m *m) UserValid() *User {
	user, err := New(&Required{
		VerifyToken: m.IDValid().Value(),
		ID:          m.IDValid().Value(),
		Email:       m.EmailValid().Value(),
		Username:    m.UsernameValid().Value(),
		Password:    m.PasswordValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

var Mother = mother.New[m]
