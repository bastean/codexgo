package authentications

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/golang-jwt/jwt/v5"
)

type Payload = map[string]any

type JWT struct {
	secretKey []byte
}

func (auth *JWT) Generate(payload map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	signature, err := token.SignedString(auth.secretKey)

	if err != nil {
		return "", errors.NewInternal(&errors.Bubble{
			Where: "Generate",
			What:  "failure to sign a jwt",
			Who:   err,
		})
	}

	return signature, nil
}

func (auth *JWT) Validate(signature string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(signature, func(token *jwt.Token) (any, error) {
		return auth.secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.NewFailure(&errors.Bubble{
		Where: "Validate",
		What:  "invalid jwt signature",
		Why: errors.Meta{
			"JWT": signature,
		},
	})
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		secretKey: []byte(secretKey),
	}
}
