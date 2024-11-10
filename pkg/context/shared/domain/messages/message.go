package messages

type Message = struct {
	ID         string
	OccurredOn string
	Key        Key
	Attributes any
	Meta       any
}

func New[T ~Message](key Key, attributes, meta any) *T {
	return &T{
		Key:        key,
		Attributes: attributes,
		Meta:       meta,
	}
}
