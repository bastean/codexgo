package errs

type InternalError struct {
	*Bubble
}

func NewInternalError(bubble *Bubble) error {
	return &InternalError{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
