package messages

type Message struct {
	ID         string
	OccurredOn string
	Key        Key
	Attributes any
	Meta       any
}

func New(key Key, attributes, meta any) *Message {
	return &Message{
		Key:        key,
		Attributes: attributes,
		Meta:       meta,
	}
}
