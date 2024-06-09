package login

type Response struct {
	Id, Email, Username, Password string
	Verified                      bool
}
