package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

var (
	smtp   = strings.Split(os.Getenv("CODEXGO_SMTP_USERNAME"), "@")
	user   = smtp[0]
	domain = smtp[1]
)

type mother struct {
	*gofakeit.Faker
}

func (m *mother) Email() string {
	random := strings.Split(m.Faker.Email(), "@")[0]

	switch {
	case len(smtp) == 2:
		return fmt.Sprintf(
			"%s+%s@%s",
			user,
			random,
			domain,
		)
	default:
		return fmt.Sprintf(
			"%s@%s",
			random,
			"example.com",
		)
	}
}

func (m *mother) Message() string {
	return m.LoremIpsumSentence(m.RandomInt([]int{1, 10}))
}

func (*mother) ID() string {
	return GenerateID()
}

var Create = &mother{
	Faker: gofakeit.New(0),
}
