package values

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type IntPrimitive = Primitive[int]

type Int struct {
	Object[int]
}

type IntPositive struct {
	Int
}

func (i *IntPositive) Validate() error {
	if i.RawValue() < 0 {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Invalid positive number",
			Why: errors.Meta{
				"Number": i.RawValue(),
			},
		})
	}

	i.Valid()

	return nil
}

type IntNegative struct {
	Int
}

func (i *IntNegative) Validate() error {
	if i.RawValue() > 0 {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Invalid negative number",
			Why: errors.Meta{
				"Number": i.RawValue(),
			},
		})
	}

	i.Valid()

	return nil
}
