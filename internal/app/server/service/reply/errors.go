package reply

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func Errors(bubbles *errors.Bubbles) []*Error {
	if bubbles.Amount == 0 {
		return nil
	}

	for _, unknown := range bubbles.Unknown {
		log.Error("Unknown | " + unknown.Error())
	}

	errors := make([]*Error, 0, (bubbles.Amount - (len(bubbles.Internal) + len(bubbles.Unknown))))

	for _, failure := range bubbles.Failure {
		errors = append(errors, &Error{Type: "Failure", Message: failure.What, Data: failure.Why})
	}

	for _, invalidValue := range bubbles.InvalidValue {
		errors = append(errors, &Error{Type: "InvalidValue", Message: invalidValue.What, Data: invalidValue.Why})
	}

	for _, alreadyExist := range bubbles.AlreadyExist {
		errors = append(errors, &Error{Type: "AlreadyExist", Message: alreadyExist.What, Data: alreadyExist.Why})
	}

	for _, notExist := range bubbles.NotExist {
		errors = append(errors, &Error{Type: "NotExist", Message: notExist.What, Data: notExist.Why})
	}

	return errors
}
