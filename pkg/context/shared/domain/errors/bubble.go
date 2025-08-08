package errors

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/id"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

const (
	Separator = "|"
)

type (
	Meta = map[string]any
)

type Bubble struct {
	ID          string
	When        time.Time
	Where, What string
	Why         Meta
	Who         error
}

type (
	Default      struct{ *Bubble }
	Internal     struct{ *Bubble }
	Failure      struct{ *Bubble }
	InvalidValue struct{ *Bubble }
	AlreadyExist struct{ *Bubble }
	NotExist     struct{ *Bubble }
)

func (b *Bubble) Error() string {
	message := fmt.Sprintf("#%s %s %s (%s): %s", b.ID, Separator, b.When.Format(), b.Where, b.What)

	if b.Why != nil {
		why, err := json.Marshal(b.Why)

		if err != nil {
			Panic(Standard("Cannot format \"Why\" from error Bubble [%s]", err))
		}

		message = fmt.Sprintf("%s %s %s", message, Separator, why)
	}

	if b.Who != nil {
		message = fmt.Sprintf("%s %s [%s]", message, Separator, b.Who)
	}

	return message
}

func New[Error ~struct{ *Bubble }](bubble *Bubble) *Error {
	bubble.ID = id.New()

	bubble.When = time.Now()

	if bubble.Where == "" {
		where, _, _, _ := caller.Received(caller.SkipCurrent)
		bubble.Where = where
	}

	if bubble.What == "" {
		Panic(Standard("Cannot create a error Bubble if \"What\" is not defined"))
	}

	return &Error{bubble}
}
