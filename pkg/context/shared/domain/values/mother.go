package values

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type m struct {
	*mother.Mother
}

func (m *m) IDNew(id string) *ID {
	value, err := New[*ID](id)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IDValid() *ID {
	value, err := New[*ID](m.UUID())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IDInvalid() (string, error) {
	value := "x"

	_, err := New[*ID](value)

	return value, err
}

func (m *m) TokenNew(token string) *Token {
	value, err := New[*Token](token)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) TokenValid() *Token {
	value, err := New[*Token](m.UUID())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) TokenInvalid() (string, error) {
	value := "x"

	_, err := New[*Token](value)

	return value, err
}

func (m *m) TimeValid() *Time {
	value, err := New[*Time](time.Now().Format())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) TimeInvalid() (string, error) {
	var value string

	_, err := New[*Time](value)

	return value, err
}

func (m *m) EmailNew(email string) *Email {
	value, err := New[*Email](email)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) EmailValid() *Email {
	value, err := New[*Email](m.Email())

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

	_, err := New[*Email](value)

	return value, err
}

func (m *m) UsernameNew(username string) *Username {
	value, err := New[*Username](username)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) UsernameValid() *Username {
	value, err := New[*Username](m.Regex(`^[A-Za-z0-9]{2,20}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) UsernameInvalidLength() (string, error) {
	value := "x"

	_, err := New[*Username](value)

	return value, err
}

func (m *m) UsernameInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := New[*Username](value)

	return value, err
}

var Mother = mother.New[m]
