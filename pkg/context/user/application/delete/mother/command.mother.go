package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *delete.Command {
	return delete.NewCommand(valueObjectMother.RandomId().Value)
}

func Invalid() *delete.Command {
	return delete.NewCommand(valueObjectMother.InvalidId().Value)
}
