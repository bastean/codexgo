package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Root struct {
	Created, Updated *Time
	Events           []*messages.Message
}

func (r *Root) CreationStamp() error {
	if r.Created != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Cannot overwrite an existing stamp",
		})
	}

	created, err := values.New[*Time](time.Now().Format())

	if err != nil {
		return errors.BubbleUp(err)
	}

	r.Created = created
	r.Updated = created

	return nil
}

func (r *Root) UpdatedStamp() error {
	updated, err := values.Replace(r.Updated, time.Now().Format())

	if err != nil {
		return errors.BubbleUp(err)
	}

	r.Updated = updated

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
