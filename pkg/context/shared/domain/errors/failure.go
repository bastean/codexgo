package errors

type Failure struct {
	*Bubble
}

func NewFailure(bubble *Bubble) error {
	return &Failure{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
