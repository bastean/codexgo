package login

type Response struct {
	Id       string
	Email    string
	Username string
	Password string
}

func NewResponse(id, email, username, password string) *Response {
	return &Response{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
	}
}
