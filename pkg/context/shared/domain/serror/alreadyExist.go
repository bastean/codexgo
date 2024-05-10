package serror

type AlreadyExistError struct {
	*Bubble
}

func NewAlreadyExistError(bubble *Bubble) error {
	return &AlreadyExistError{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
