package smessage

type Message struct {
	Id         string
	Type       string
	OccurredOn string
	Attributes []byte
	Meta       []byte
}

func NewMessage(routingKey string, attributes, meta []byte) *Message {
	return &Message{
		Type:       routingKey,
		Attributes: attributes,
		Meta:       meta,
	}
}
