package env

import (
	"os"
)

var (
	JWTSecretKey string
)

func JWT() {
	JWTSecretKey = os.Getenv(JWT_SECRET_KEY)
}
