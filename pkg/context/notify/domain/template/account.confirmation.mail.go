package template

type AccountConfirmationMail struct {
	*Mail
	Username         string
	ConfirmationLink string
}
