package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

var (
	smtp = strings.Split(os.Getenv("CODEXGO_SMTP_USERNAME"), "@")
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
			smtp[0],
			random,
			smtp[1],
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
	return m.LoremIpsumSentence(m.IntRange(1, 12))
}

func (m *mother) TimeSetBefore(actual time.Time, min, max int) time.Time {
	return actual.Add(-time.Duration(m.IntRange(min, max)))
}

func (m *mother) TimeRandomBefore(actual time.Time) time.Time {
	return m.TimeSetBefore(actual, 1, 72)
}

func (m *mother) TimeNow() time.Time {
	return time.Now()
}

func (m *mother) TimeSetAfter(actual time.Time, min, max int) time.Time {
	return actual.Add(time.Duration(m.IntRange(min, max)))
}

func (m *mother) TimeRandomAfter(actual time.Time) time.Time {
	return m.TimeSetAfter(actual, 1, 72)
}

func (*mother) ID() string {
	return GenerateID()
}

var Create = &mother{
	Faker: gofakeit.New(0),
}
