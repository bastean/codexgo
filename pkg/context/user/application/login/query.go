package login

type Query struct {
	Email    string
	Password string
}

func NewQuery(email, password string) *Query {
	return &Query{
		Email:    email,
		Password: password,
	}
}
