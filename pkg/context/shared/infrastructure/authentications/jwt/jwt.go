package jwt

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	Payload = map[string]any
)

type JWT struct {
	secretKey []byte
}

func (j *JWT) Generate(payload map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	signature, err := token.SignedString(j.secretKey)

	if err != nil {
		return "", errors.New[errors.Internal](&errors.Bubble{
			Where: "Generate",
			What:  "Failure to sign a JWT",
			Who:   err,
		})
	}

	return signature, nil
}

func (j *JWT) Validate(signature string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(signature, func(token *jwt.Token) (any, error) {
		return j.secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New[errors.Failure](&errors.Bubble{
		Where: "Validate",
		What:  "Invalid JWT signature",
		Why: errors.Meta{
			"Signature": signature,
		},
	})
}

func New(secretKey string) *JWT {
	return &JWT{
		secretKey: []byte(secretKey),
	}
}
