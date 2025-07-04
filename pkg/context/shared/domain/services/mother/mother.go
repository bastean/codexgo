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

func (m *Mother) Words(amount int) []string {
	words := make([]string, amount)

	for i := range amount {
		words[i] = m.LoremIpsumWord()
	}

	return words
}

func (m *Mother) WordsJoin(words []string, sep string) string {
	return strings.Join(words, sep)
}

func (m *Mother) Message() string {
	return m.WordsJoin(m.Words(m.IntRange(1, 12)), " ")
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
