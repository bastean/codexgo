package root

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Root struct {
	CreatedAt, UpdatedAt *values.Time
	Events               []*messages.Message
}

func (r *Root) CreationStamp() error {
	if r.CreatedAt != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Cannot overwrite an existing stamp",
		})
	}

	createdAt, err := values.New[*values.Time](time.Now().Format())

	if err != nil {
		return errors.BubbleUp(err)
	}

	r.CreatedAt = createdAt

	return nil
}

func (r *Root) UpdatedStamp() error {
	var (
		err       error
		updatedAt *values.Time
	)

	switch r.UpdatedAt {
	case nil:
		updatedAt, err = values.New[*values.Time](time.Now().Format())
	default:
		updatedAt, err = values.Replace(r.UpdatedAt, time.Now().Format())
	}

	if err != nil {
		return errors.BubbleUp(err)
	}

	r.UpdatedAt = updatedAt

	return nil
}

func (r *Root) Record(events ...*messages.Message) {
	r.Events = append(r.Events, events...)
}

func (r *Root) Pull() []*messages.Message {
	recorded := r.Events

	r.Events = []*messages.Message{}

	return recorded
}
