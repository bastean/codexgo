package errors

type AlreadyExist struct {
	*Bubble
}

func NewAlreadyExist(bubble *Bubble) error {
	return &AlreadyExist{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
