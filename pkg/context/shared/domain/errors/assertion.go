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

func EventAssertion(where string) error {
	return assertion("Event", where)
}

func CommandAssertion(where string) error {
	return assertion("Command", where)
}

func QueryAssertion(where string) error {
	return assertion("Query", where)
}
