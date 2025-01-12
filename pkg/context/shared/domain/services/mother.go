package services

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/ids"
)

type mother struct {
	*gofakeit.Faker
}

func (m *mother) Email() string {
	username := strings.Split(m.Faker.Email(), "@")[0]

	domain := "example.com"

	return fmt.Sprintf("%s@%s", username, domain)
}

func (*mother) ID() string {
	return ids.Generate()
}

var Create = &mother{
	Faker: gofakeit.New(0),
}
