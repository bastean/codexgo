package commands

type (
	Type string
)

type Command interface {
	Type() Type
}

type Handler interface {
	SubscribedTo() Type
	Handle(Command) error
}

type Bus interface {
	Register(Type, Handler) error
	Dispatch(Command) error
}
