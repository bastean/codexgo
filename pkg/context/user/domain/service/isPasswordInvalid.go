package service

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

func IsPasswordInvalid(hashing model.Hashing, hashed, plain string) error {
	if hashing.IsNotEqual(hashed, plain) {
		return errors.NewFailure(&errors.Bubble{
			Where: "IsPasswordInvalid",
			What:  "passwords do not match",
		})
	}

	return nil
}
