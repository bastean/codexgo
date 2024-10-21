package query

type (
	Type string
)

type Query interface {
	Type() Type
}

type Response interface {
	Type() Type
}

type Handler interface {
	SubscribedTo() Type
	ReplyTo() Type
	Handle(Query) (Response, error)
}

type Bus interface {
	Register(Type, Handler) error
	Ask(Query) (Response, error)
}
