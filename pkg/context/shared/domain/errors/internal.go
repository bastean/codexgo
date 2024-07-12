package errors

type ErrInternal struct {
	*Bubble
}

func NewInternal(bubble *Bubble) error {
	return &ErrInternal{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
