package services

import (
	"errors"

	"github.com/bastean/codexgo/context/pkg/user/domain/models"
)

func IsPasswordInvalid(hashing models.Hashing, plain, hashed string) (err error) {
	if hashing.IsNotEqual(plain, hashed) {
		err = errors.New("incorrect password")
	}

	return
}
