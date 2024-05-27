package errors

type AlreadyExist struct {
	*Bubble
}

type NotExist struct {
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

func NewNotExist(bubble *Bubble) error {
	return &NotExist{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
