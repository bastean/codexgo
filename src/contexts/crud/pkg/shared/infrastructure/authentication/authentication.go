package authentication

import (
	"os"
	"time"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/golang-jwt/jwt"
)

var InvalidJWT = errors.InvalidValue{Dump: []errors.Error{{Field: "JWT", Message: "Invalid"}}}

var secretKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

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

func ValidateJWT(tokenString string) jwt.MapClaims {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}

	panic(InvalidJWT)
}
