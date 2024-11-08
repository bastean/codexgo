package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type Root struct {
	Events []*events.Event
}

func (r *Root) Record(events ...*events.Event) {
	r.Events = append(r.Events, events...)
}

func (r *Root) Pull() []*events.Event {
	recorded := r.Events

	r.Events = []*events.Event{}

	return recorded
}

func NewRoot() *Root {
	return &Root{
		Events: []*events.Event{},
	}
}
