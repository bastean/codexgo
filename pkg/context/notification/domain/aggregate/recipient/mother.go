package recipient

import (
	"strings"

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

func (m *m) RecipientValid() *Recipient {
	user, err := New(&Required{
		ID:       m.IDValid().Value(),
		Email:    m.EmailValid().Value(),
		Username: m.UsernameValid().Value(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return user
}

var Mother = mother.New[m]
