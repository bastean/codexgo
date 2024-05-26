package authentications

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/golang-jwt/jwt"
)

type Payload map[string]any

type Authentication struct {
	secretKey []byte
}

func (auth *Authentication) GenerateJWT(payload map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	tokenString, err := token.SignedString(auth.secretKey)

	if err != nil {
		return "", errors.NewInternal(&errors.Bubble{
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

	return nil, errors.NewFailure(&errors.Bubble{
		Where: "ValidateJWT",
		What:  "invalid jwt signature",
		Why: errors.Meta{
			"JWT": tokenString,
		},
	})
}

func NewAuthentication(secretKey string) *Authentication {
	return &Authentication{
		secretKey: []byte(secretKey),
	}
}
