package update

type Command struct {
	Id              string
	Email           string
	Username        string
	Password        string
	UpdatedPassword string
}

func NewCommand(id, email, username, password, updatedPassword string) *Command {
	return &Command{
		Id:              id,
		Email:           email,
		Username:        username,
		Password:        password,
		UpdatedPassword: updatedPassword,
	}
}
