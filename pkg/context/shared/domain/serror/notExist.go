package serror

type NotExistError struct {
	*Bubble
}

func NewNotExistError(bubble *Bubble) error {
	return &NotExistError{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
