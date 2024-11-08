package events

type (
	Key       string
	Recipient string
)

type Event struct {
	ID         string
	OccurredOn string
	Key        Key
	Attributes any
	Meta       any
}

type Consumer interface {
	On(*Event) error
}

type Bus interface {
	Subscribe(Key, Consumer) error
	Publish(*Event) error
}

func New(key Key, attributes, meta any) *Event {
	return &Event{
		Key:        key,
		Attributes: attributes,
		Meta:       meta,
	}
}
