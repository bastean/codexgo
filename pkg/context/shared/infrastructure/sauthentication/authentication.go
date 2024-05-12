package sauthentication

import (
	"time"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/golang-jwt/jwt"
)

type Authentication struct {
	secretKey []byte
}

func (auth *Authentication) GenerateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add((24 * time.Hour) * 7).Unix(),
		"userId": userId,
	})

	tokenString, err := token.SignedString(auth.secretKey)

	if err != nil {
		return "", serror.NewInternal(&serror.Bubble{
			Where: "GenerateJWT",
			What:  "failure to sign a jwt",
			Who:   err,
		})
	}

	return tokenString, nil
}

func (auth *Authentication) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return auth.secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, serror.NewFailure(&serror.Bubble{
		Where: "ValidateJWT",
		What:  "invalid jwt signature",
		Why: serror.Meta{
			"JWT": tokenString,
		},
	})
}

func NewAuthentication(secretKey string) *Authentication {
	return &Authentication{
		secretKey: []byte(secretKey),
	}
}
