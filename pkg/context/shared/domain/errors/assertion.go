package errors

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func assertion(what string) error {
	_, _, where := caller.Received(caller.SkipCurrent + 1)

	return New[Internal](&Bubble{
		Where: where,
		What:  fmt.Sprintf("Failure in %s type assertion", what),
	})
}

func EventAssertion() error {
	return assertion("Event")
}

func CommandAssertion() error {
	return assertion("Command")
}

func QueryAssertion() error {
	return assertion("Query")
}

func IsNot(err, target error) bool {
	return err != nil && !Is(err, target)
}
