package smodel

type CommandHandler[Command any] interface {
	Handle(Command) error
}
