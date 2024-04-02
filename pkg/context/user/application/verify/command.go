package verify

type Command struct {
	Id string
}

func NewCommand(id string) *Command {
	return &Command{
		Id: id,
	}
}
