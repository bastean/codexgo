package update

type CommandHandler struct {
	*Update
}

func (commandHandler *CommandHandler) Handle(command *Command) {
	commandHandler.Update.Run(command)
}

func NewCommandHandler(update *Update) *CommandHandler {
	return &CommandHandler{
		update,
	}
}
