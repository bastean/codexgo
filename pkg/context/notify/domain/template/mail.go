package template

type Mail struct {
	To []string
}

func NewMail(to []string) *Mail {
	return &Mail{
		To: to,
	}
}
