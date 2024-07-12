package errors

type ErrFailure struct {
	*Bubble
}

func NewFailure(bubble *Bubble) error {
	return &ErrFailure{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
