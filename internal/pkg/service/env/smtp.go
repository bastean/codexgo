package env

import (
	"os"
)

var (
	SMTPHost, SMTPPort, SMTPUsername, SMTPPassword string
)

func SMTP() {
	SMTPHost = os.Getenv(SMTP_HOST)
	SMTPPort = os.Getenv(SMTP_PORT)
	SMTPUsername = os.Getenv(SMTP_USERNAME)
	SMTPPassword = os.Getenv(SMTP_PASSWORD)
}

func HasSMTP() bool {
	return SMTPHost != ""
}
