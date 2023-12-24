package services

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
)

var IncorrectPassword = errors.InvalidValue{Dump: []errors.Error{{Field: "Password", Message: "Incorrect"}}}

func IsPasswordInvalid(hashing models.Hashing, hashed, plain string) {
	if hashing.IsNotEqual(hashed, plain) {
		panic(IncorrectPassword)
	}
}
