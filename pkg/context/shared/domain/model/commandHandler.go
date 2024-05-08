package model

type CommandHandler[Command any] interface {
	Handle(Command) error
}
