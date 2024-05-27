package errors

type InvalidValue struct {
	*Bubble
}

func NewInvalidValue(bubble *Bubble) error {
	return &InvalidValue{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
