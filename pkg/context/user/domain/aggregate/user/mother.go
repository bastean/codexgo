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

func (m *m) CipherPasswordValid() *CipherPassword {
	value, err := values.New[*CipherPassword](m.Regex(`^.{8,64}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) CipherPasswordInvalid() (string, error) {
	var value string

	_, err := values.New[*CipherPassword](value)

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

func (m *m) UserValidPrimitive() *User {
	user, err := FromPrimitive(&Primitive{
		Created:  aggregates.Mother.TimeValid().Value(),
		Updated:  aggregates.Mother.TimeValid().Value(),
		Verify:   m.IDValid().Value(),
		Reset:    m.IDValid().Value(),
		ID:       m.IDValid().Value(),
		Email:    m.EmailValid().Value(),
		Username: m.UsernameValid().Value(),
		Password: m.PlainPasswordValid().Value(),
		Verified: m.VerifiedValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

func (m *m) UserValidRaw() *User {
	user, err := FromRaw(&Primitive{
		ID:       m.IDValid().Value(),
		Email:    m.EmailValid().Value(),
		Username: m.UsernameValid().Value(),
		Password: m.PlainPasswordValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

func (m *m) UserValid() *User {
	user := m.UserValidPrimitive()

	user.Created = nil
	user.Updated = nil
	user.Reset = nil

	return user
}

var Mother = mother.New[m]()
