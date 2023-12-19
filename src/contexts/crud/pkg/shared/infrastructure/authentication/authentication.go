package authentication

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("codexgo") //! .env

func GenerateJWT(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add((24 * time.Hour) * 7).Unix(),
		"id":  id,
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		panic(err.Error())
	}

	return tokenString
}

func ValidateJWT(tokenString string) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if !token.Valid {
		panic("invalid JWT")
	}
}
