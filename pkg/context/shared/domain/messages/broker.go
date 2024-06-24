package messages

type Broker interface {
	AddRouter(router *Router) error
	AddQueue(queue *Queue) error
	AddQueueMessageBind(queue *Queue, bindingKeys []string) error
	AddQueueConsumer(consumer Consumer) error
	PublishMessages(messages []*Message) error
}
