package errors

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bubble struct {
	When        time.Time
	Where, What string
	Why         map[string]any
	Who         error
}

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

func NewBubble(where, what string, why Meta, who error) *Bubble {
	if where == "" {
		Panic("Cannot create a error Bubble if \"Where\" is not defined", "NewBubble")
	}

	if what == "" {
		Panic("Cannot create a error Bubble if \"What\" is not defined", "NewBubble")
	}

	return &Bubble{
		When:  time.Now().UTC(),
		Where: where,
		What:  what,
		Why:   why,
		Who:   who,
	}
}
