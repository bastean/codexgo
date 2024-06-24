package messages

type Consumer interface {
	SubscribedTo() []*Queue
	On(message *Message) error
}
