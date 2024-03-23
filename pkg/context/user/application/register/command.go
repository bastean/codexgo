package register

type Command struct {
	Id       string
	Email    string
	Username string
	Password string
}

func NewCommand(id, email, username, password string) *Command {
	return &Command{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
	}
}
