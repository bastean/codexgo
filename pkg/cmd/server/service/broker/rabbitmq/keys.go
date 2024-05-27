package rabbitmq

type EventKey struct {
	CreatedSucceeded string
}

type Key struct {
	Event *EventKey
}

var BindingKey = &Key{
	Event: &EventKey{
		CreatedSucceeded: "#.event.#.created.succeeded",
	},
}
