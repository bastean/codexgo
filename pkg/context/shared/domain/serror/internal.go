package serror

type Internal struct {
	*Bubble
}

func NewInternal(bubble *Bubble) error {
	return &Internal{
		Bubble: NewBubble(
			bubble.Where,
			bubble.What,
			bubble.Why,
			bubble.Who,
		),
	}
}
