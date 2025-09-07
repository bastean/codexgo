package mother

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/id"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

var (
	smtp = strings.Split(os.Getenv("CODEXGO_SMTP_USERNAME"), "@")
)

type Mother struct {
	*gofakeit.Faker
}

func Generator[T any](amount int, value func() T) []T {
	values := make([]T, amount)

	for i := range amount {
		values[i] = value()
	}

	return values
}

func (m *Mother) Email() string {
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

func (m *Mother) Letters(amount int) []string {
	return Generator(amount, m.Letter)
}

func (m *Mother) Words(amount int) []string {
	return Generator(amount, m.LoremIpsumWord)
}

func (m *Mother) Join(values []string, separator string) string {
	return strings.Join(values, separator)
}

func (m *Mother) Message() string {
	return m.Join(m.Words(m.IntRange(1, 12)), " ")
}

func (m *Mother) TimeSetBefore(actual time.Time, min, max time.Duration) time.Time {
	return actual.Add(-time.Duration(m.IntRange(int(min), int(max))))
}

func (m *Mother) TimeRandomBefore(actual time.Time) time.Time {
	return m.TimeSetBefore(actual, time.Day, time.Day*3)
}

func (m *Mother) TimeNow() time.Time {
	return time.Now()
}

func (m *Mother) TimeSetAfter(actual time.Time, min, max time.Duration) time.Time {
	return actual.Add(time.Duration(m.IntRange(int(min), int(max))))
}

func (m *Mother) TimeRandomAfter(actual time.Time) time.Time {
	return m.TimeSetAfter(actual, time.Day, time.Day*3)
}

func (m *Mother) StructRandomize(value any) {
	if err := m.Struct(value); err != nil {
		log.Panic(err)
	}
}

func (*Mother) ID() string {
	return id.New()
}

func New[M ~struct{ *Mother }]() *M {
	if _, ok := os.LookupEnv("GOTEST"); !ok {
		log.Panic("Use \"Mother\" only in a \"Test Environment\"")
	}

	return &M{
		Mother: &Mother{
			Faker: gofakeit.New(0),
		},
	}
}
