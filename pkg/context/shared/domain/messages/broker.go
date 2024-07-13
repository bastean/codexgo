package messages

type Broker interface {
	AddRouter(*Router) error
	AddQueue(*Queue) error
	AddQueueMessageBind(*Queue, BindingKeys) error
	AddQueueConsumer(Consumer) error
	PublishMessages([]*Message) error
}
