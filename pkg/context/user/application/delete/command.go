package delete

type Command struct {
	Id string
}

func NewCommand(id string) *Command {
	return &Command{
		Id: id,
	}
}
