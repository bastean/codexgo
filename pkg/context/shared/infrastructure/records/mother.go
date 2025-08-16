package records

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) Message() (format string, values []any, message string) {
	format = "%s %d %s"

	values = []any{m.LoremIpsumWord(), m.Int(), m.LoremIpsumWord()}

	return format, values, fmt.Sprintf(format, values...)
}

var Mother = mother.New[m]
