package errs

type InvalidValueError struct {
	*Bubble
}

func NewInvalidValueError(bubble *Bubble) error {
	return &InvalidValueError{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
