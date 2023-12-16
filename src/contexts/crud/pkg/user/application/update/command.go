package update

type Command struct {
	Id              string
	Email           string
	Username        string
	CurrentPassword string
	UpdatedPassword string
}
