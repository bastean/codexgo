package models

type CommandHandler[Command any] interface {
	Handle(Command) error
}

type QueryHandler[Query, Response any] interface {
	Handle(Query) (Response, error)
}
