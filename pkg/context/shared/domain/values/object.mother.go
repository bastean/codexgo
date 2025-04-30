package values

import (
	"os"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type Example struct {
	Object[string]
}

func (m *Example) Validate() error {
	if m.RawValue() == "" {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Value can not be empty",
		})
	}

	m.Valid()

	return nil
}

type mother struct{}

func (m *mother) ObjectWithCustomValidation() *Example {
	return new(Example)
}

func Mother() *mother {
	if _, ok := os.LookupEnv("GOTEST"); !ok {
		errors.Panic(errors.Standard("Use only in a \"Test Environment\""))
	}

	return new(mother)
}
