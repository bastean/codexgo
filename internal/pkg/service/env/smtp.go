package env

import (
	"os"
)

var (
	SMTPHost     = os.Getenv("CODEXGO_SMTP_HOST")
	SMTPPort     = os.Getenv("CODEXGO_SMTP_PORT")
	SMTPUsername = os.Getenv("CODEXGO_SMTP_USERNAME")
	SMTPPassword = os.Getenv("CODEXGO_SMTP_PASSWORD")
)
