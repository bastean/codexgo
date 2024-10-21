package errors

import (
	"fmt"
)

func assertion(what, where string) error {
	return New[Internal](&Bubble{
		Where: where,
		What:  fmt.Sprintf("Failure in %s type assertion", what),
	})
}

func CommandAssertion(where string) error {
	return assertion("Command", where)
}

func QueryAssertion(where string) error {
	return assertion("Query", where)
}
