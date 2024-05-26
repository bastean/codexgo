package errors

type NotExist struct {
	*Bubble
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
