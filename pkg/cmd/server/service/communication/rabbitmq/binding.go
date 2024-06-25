package rabbitmq

type Event struct {
	CreatedSucceeded string
}

type Key struct {
	*Event
}

var Binding = &Key{
	Event: &Event{
		CreatedSucceeded: "#.event.#.created.succeeded",
	},
}
