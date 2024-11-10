package messages

type Message = struct {
	ID         string
	OccurredOn string
	Key        Key
	Attributes any
	Meta       any
}
