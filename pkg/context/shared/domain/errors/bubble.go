package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

var (
	Join = errors.Join
	As   = errors.As
	Is   = errors.Is
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
	Internal     struct{ *Bubble }
	Failure      struct{ *Bubble }
	InvalidValue struct{ *Bubble }
	AlreadyExist struct{ *Bubble }
	NotExist     struct{ *Bubble }
)

type Error interface {
	Internal | Failure | InvalidValue | AlreadyExist | NotExist
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

func New[Err Error](bubble *Bubble) *Err {
	return &Err{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}

func Default() error {
	return new(Bubble)
}

func BubbleUp(who error, where string) error {
	return fmt.Errorf("(%s): [%w]", where, who)
}

func Panic(what, where string) {
	log.Panicf("(%s): %s", where, what)
}
