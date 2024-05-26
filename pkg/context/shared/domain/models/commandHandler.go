package models

type CommandHandler[Command any] interface {
	Handle(Command) error
}
