package messages

type Consumer interface {
	SubscribedTo() []*Queue
	On(*Message) error
}
