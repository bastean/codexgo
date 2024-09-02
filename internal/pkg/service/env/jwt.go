package env

import (
	"os"
)

var (
	JWTSecretKey string
)

func JWT() {
	JWTSecretKey = os.Getenv("CODEXGO_JWT_SECRET_KEY")
}
