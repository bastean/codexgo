package services

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

type mother struct {
	*gofakeit.Faker
}

func (m *mother) Email() string {
	username := strings.Split(m.Faker.Email(), "@")[0]

	domain := "example.com"

	return fmt.Sprintf("%s@%s", username, domain)
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
