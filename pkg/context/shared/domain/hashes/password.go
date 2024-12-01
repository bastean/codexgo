package hashes

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func IsPasswordInvalid(hasher Hasher, hashed, plain string) error {
	if hasher.IsNotEqual(hashed, plain) {
		return errors.New[errors.Failure](&errors.Bubble{
			Where: "IsPasswordInvalid",
			What:  "Invalid password",
		})
	}

	return nil
}
