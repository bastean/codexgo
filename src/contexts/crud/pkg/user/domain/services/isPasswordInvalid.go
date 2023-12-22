package services

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
)

var IncorrectPassword = errors.InvalidValue{Message: "incorrect password"}

func IsPasswordInvalid(hashing models.Hashing, hashed, plain string) {
	if hashing.IsNotEqual(hashed, plain) {
		panic(IncorrectPassword)
	}
}
