package attempt

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) AttemptCopy(attempt *Attempt) *Attempt {
	copy, err := FromPrimitive(attempt.ToPrimitive())

	if err != nil {
		errors.Panic(err)
	}

	return copy
}

func (m *m) AttemptRequiredValid() *Required {
	return &Required{
		Limit: values.Mother().IntPositiveValid().Value(),
		Every: values.Mother().IntPositiveValid().Value(),
		Next:  values.Mother().IntPositiveValid().Value(),
	}
}

func (m *m) AttemptValid() *Attempt {
	attempt, err := New(m.AttemptRequiredValid())

	if err != nil {
		errors.Panic(err)
	}

	return attempt
}

func (m *m) AttemptValidFromPrimitive(without ...string) *Attempt {
	attempt, err := FromPrimitive(&Primitive{
		Counter: values.Mother().IntPositiveValid().ToPrimitive(),
		Limit:   values.Mother().IntPositiveValid().ToPrimitive(),
		Until:   values.Mother().TimeValid().ToPrimitive(),
		Every:   values.Mother().IntPositiveValid().ToPrimitive(),
		Next:    values.Mother().IntPositiveValid().ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err)
	}

	for _, field := range without {
		switch field {
		case "Until":
			attempt.Until = nil
		default:
			errors.Panic(errors.Standard("Unknown field %q", field))
		}
	}

	return attempt
}

var Mother = mother.New[m]
