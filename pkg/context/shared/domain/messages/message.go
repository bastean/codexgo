package messages

type Message struct {
	Id, Type, OccurredOn string
	Attributes, Meta     []byte
}

func NewMessage(routingKey string, attributes, meta []byte) *Message {
	return &Message{
		Type:       routingKey,
		Attributes: attributes,
		Meta:       meta,
	}
}
