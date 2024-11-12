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

var Create = &mother{
	Faker: gofakeit.New(0),
}
