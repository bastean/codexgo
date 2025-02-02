package errors

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type (
	Meta = map[string]any
)

type Bubble struct {
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
	message := fmt.Sprintf("%s (%s): %s", services.FormatTime(b.When), b.Where, b.What)

	if b.Why != nil {
		why, err := json.Marshal(b.Why)

		if err != nil {
			why = []byte(fmt.Sprintf("{\"Error\":\"Cannot JSON encoding \"Why\" from error Bubble: [%s]\"}", err))
		}

		message = fmt.Sprintf("%s: %s", message, why)
	}

	if b.Who != nil {
		message = fmt.Sprintf("%s: [%s]", message, b.Who)
	}

	return message
}

func New[Error ~struct{ *Bubble }](bubble *Bubble) *Error {
	if bubble.When.IsZero() {
		bubble.When = time.Now().UTC()
	}

	if bubble.Where == "" {
		Panic("Cannot create a error Bubble if \"Where\" is not defined", "NewBubble")
	}

	if bubble.What == "" {
		Panic("Cannot create a error Bubble if \"What\" is not defined", "NewBubble")
	}

	return &Error{bubble}
}

func BubbleUp(who error, where string) error {
	return fmt.Errorf("(%s): [%w]", where, who)
}

func IsNot(err error, target error) bool {
	return err != nil && !Is(err, target)
}
