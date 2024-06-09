package errors

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bubble struct {
	When  time.Time
	Where string
	What  string
	Why   map[string]any
	Who   error
}

func (err *Bubble) Error() string {
	message := fmt.Sprintf("%s (%s): %s", err.When.Format(time.RFC3339Nano), err.Where, err.What)

	if err.Why != nil {
		why, err := json.Marshal(err.Why)

		if err != nil {
			PanicOnError("Error", fmt.Sprintf("cannot json encoding why from error bubble: %s: [%s]", message, err.Error()))
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
		PanicOnError("NewBubble", "cannot create a error bubble if where is not defined")
	}

	if what == "" {
		PanicOnError("NewBubble", "cannot create a error bubble if what is not defined")
	}

	return &Bubble{
		When:  time.Now().UTC(),
		Where: where,
		What:  what,
		Why:   why,
		Who:   who,
	}
}
