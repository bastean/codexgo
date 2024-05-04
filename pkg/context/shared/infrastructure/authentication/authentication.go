package authentication

import (
	"time"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/golang-jwt/jwt"
)

var InvalidJWT = errors.NewInvalidValue("JWT Invalid")

type Authentication struct {
	secretKey []byte
}

func (auth *Authentication) GenerateJWT(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add((24 * time.Hour) * 7).Unix(),
		"id":  id,
	})

	tokenString, err := token.SignedString(auth.secretKey)

	if err != nil {
		panic(err.Error())
	}

	return tokenString
}

func (auth *Authentication) ValidateJWT(tokenString string) jwt.MapClaims {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return auth.secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}

	panic(InvalidJWT)
}

func NewAuthentication(secretKey string) *Authentication {
	return &Authentication{
		secretKey: []byte(secretKey),
	}
}
