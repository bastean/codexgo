package errors

import (
	"encoding/json"
	"fmt"
	"time"
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

func (err *Bubble) Error() string {
	message := fmt.Sprintf("%s (%s): %s", err.When.Format(time.RFC3339Nano), err.Where, err.What)

	if err.Why != nil {
		why, err := json.Marshal(err.Why)

		if err != nil {
			why = []byte(fmt.Sprintf("{\"Error\":\"Cannot JSON encoding \"Why\" from error Bubble: [%s]\"}", err))
		}

		message = fmt.Sprintf("%s: %s", message, why)
	}

	if err.Who != nil {
		message = fmt.Sprintf("%s: [%s]", message, err.Who)
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
