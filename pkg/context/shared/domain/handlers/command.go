package handlers

type Command[Command any] interface {
	Handle(Command) error
}
