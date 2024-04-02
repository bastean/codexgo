package sendMail

type RegisteredSucceededEventAttributes struct {
	Id       string
	Email    string
	Username string
}

func NewRegisteredSucceededEventAttributes(id, email, username string) *RegisteredSucceededEventAttributes {
	return &RegisteredSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	}
}
