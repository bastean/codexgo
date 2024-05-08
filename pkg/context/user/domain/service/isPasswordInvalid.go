package service

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

func IsPasswordInvalid(hashing model.Hashing, hashed, plain string) error {
	if hashing.IsNotEqual(hashed, plain) {
		return errs.NewFailedError(&errs.Bubble{
			Where: "IsPasswordInvalid",
			What:  "passwords do not match",
		})
	}

	return nil
}
