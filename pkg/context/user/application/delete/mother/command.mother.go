package commandMother

import (
	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
)

func Random() *delete.Command {
	return delete.NewCommand(sharedValueObjectMother.RandomId().Value)
}

func Invalid() *delete.Command {
	return delete.NewCommand(sharedValueObjectMother.InvalidId().Value)
}
