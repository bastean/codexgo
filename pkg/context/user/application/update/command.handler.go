package update

type CommandHandler struct {
	*Update
}

func (handler *CommandHandler) Handle(command *Command) {
	handler.Update.Run(command)
}

func NewCommandHandler(update *Update) *CommandHandler {
	return &CommandHandler{
		Update: update,
	}
}
