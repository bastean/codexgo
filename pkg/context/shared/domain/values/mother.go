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

func (m *m) IntNew(number int) *Int {
	value, err := New[*Int](number)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntValid() *Int {
	value, err := New[*Int](m.IntRange(-12, 12))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntPositiveNew(number int) *IntPositive {
	value, err := New[*IntPositive](number)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntPositiveValid() *IntPositive {
	value, err := New[*IntPositive](m.IntRange(1, 12))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntPositiveInvalid() (int, error) {
	value := m.IntNegativeValid().Value()

	_, err := New[*IntPositive](value)

	return value, err
}

func (m *m) IntNegativeNew(number int) *IntNegative {
	value, err := New[*IntNegative](number)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntNegativeValid() *IntNegative {
	value, err := New[*IntNegative](m.IntRange(-12, -1))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func (m *m) IntNegativeInvalid() (int, error) {
	value := m.IntPositiveValid().Value()

	_, err := New[*IntNegative](value)

	return value, err
}

func (m *m) TimeNew(time string) *Time {
	value, err := New[*Time](time)

	if err != nil {
		errors.Panic(err)
	}

	return value
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

func (m *m) EmailReplace(old *Email, value string) *Email {
	email, err := Replace(old, value)

	if err != nil {
		errors.Panic(err)
	}

	return email
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

func (m *m) UsernameReplace(old *Username, value string) *Username {
	username, err := Replace(old, value)

	if err != nil {
		errors.Panic(err)
	}

	return username
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
