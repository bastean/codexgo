package errs

type FailedError struct {
	*Bubble
}

func NewFailedError(bubble *Bubble) error {
	return &FailedError{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
