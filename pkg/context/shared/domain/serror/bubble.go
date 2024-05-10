package serror

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

type Meta map[string]string

type Bubble struct {
	When  time.Time
	Where string
	What  string
	Why   Meta
	Who   error
}

func (err *Bubble) Error() string {
	message := fmt.Sprintf("%s (%s): %s", err.When.Format(time.DateTime), err.Where, err.What)

	if err.Why != nil {
		why, _ := json.Marshal(err.Why)
		message = fmt.Sprintf("%s: %s", message, why)
	}

	if err.Who != nil {
		message = fmt.Sprintf("%s: [%s]", message, err.Who)
	}

	return message
}

func NewBubble(where, what string, why Meta, who error) *Bubble {
	if where == "" {
		sservice.PanicOnError("NewBubble", "cannot create a bubble if where is not defined")
	}

	if what == "" {
		sservice.PanicOnError("NewBubble", "cannot create a bubble if what is not defined")
	}

	return &Bubble{
		When:  time.Now(),
		Where: where,
		What:  what,
		Why:   why,
		Who:   who,
	}
}
