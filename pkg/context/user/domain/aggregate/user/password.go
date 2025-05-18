package user

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

const (
	PlainPasswordMinCharactersLength = "8"
	PlainPasswordMaxCharactersLength = "64"
)

type PlainPassword struct {
	values.Object[string]
}

func (p *PlainPassword) Validate() error {
	if values.IsNotValid(p.RawValue(), "gte=8", "lte=64") {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: fmt.Sprintf("Password must be between %s to %s characters", PlainPasswordMinCharactersLength, PlainPasswordMaxCharactersLength),
			Why: errors.Meta{
				"Password": p.RawValue(),
			},
		})
	}

	p.Valid()

	return nil
}

type Password struct {
	values.Object[string]
}

func (c *Password) Validate() error {
	if values.IsNotValid(c.RawValue(), "required") {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Password is required",
		})
	}

	c.Valid()

	return nil
}
