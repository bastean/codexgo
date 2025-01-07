package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type Root struct {
	Created, Updated *Time
	Events           []*messages.Message
}

func (r *Root) CreationStamp() error {
	if r.Created != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "CreationStamp",
			What:  "Cannot overwrite an existing stamp",
		})
	}

	created, err := NewTime(services.TimeNow())

	if err != nil {
		return errors.BubbleUp(err, "CreationStamp")
	}

	r.Created = created
	r.Updated = created

	return nil
}

func (r *Root) UpdatedStamp() error {
	updated, err := NewTime(services.TimeNow())

	if err != nil {
		return errors.BubbleUp(err, "UpdatedStamp")
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

func NewRoot() (*Root, error) {
	return &Root{
		Events: []*messages.Message{},
	}, nil
}
