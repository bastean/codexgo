package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *verify.Command {
	return verify.NewCommand(valueObjectMother.RandomId().Value)
}

func Invalid() *verify.Command {
	return verify.NewCommand(valueObjectMother.InvalidId().Value)
}
