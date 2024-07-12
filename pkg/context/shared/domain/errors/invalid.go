package errors

type ErrInvalidValue struct {
	*Bubble
}

func NewInvalidValue(bubble *Bubble) error {
	return &ErrInvalidValue{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
