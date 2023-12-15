package authentication

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("codexgo") //! .env

func GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add((24 * time.Hour) * 7),
		"id":  id,
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid JWT")
	}

	return nil
}
