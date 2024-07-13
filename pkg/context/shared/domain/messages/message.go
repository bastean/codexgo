package messages

type Attributes = []byte

type Meta = []byte

type Message struct {
	Id, Type, OccurredOn string
	Attributes, Meta     []byte
}

func New(routingKey string, attributes, meta []byte) *Message {
	return &Message{
		Type:       routingKey,
		Attributes: attributes,
		Meta:       meta,
	}
}
