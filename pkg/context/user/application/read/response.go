package read

type Response struct {
	Id       string
	Email    string
	Username string
	Password string
	Verified bool
}

func NewResponse(id, email, username, password string, verified bool) *Response {
	return &Response{
		Id:       id,
		Email:    email,
		Username: username,
		Password: password,
		Verified: verified,
	}
}
