package env

import (
	"os"
)

var (
	JWTSecretKey = os.Getenv("CODEXGO_JWT_SECRET_KEY")
)
