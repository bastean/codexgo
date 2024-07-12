package errors

type ErrAlreadyExist struct {
	*Bubble
}

type ErrNotExist struct {
	*Bubble
}

func NewAlreadyExist(bubble *Bubble) error {
	return &ErrAlreadyExist{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}

func NewNotExist(bubble *Bubble) error {
	return &ErrNotExist{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
