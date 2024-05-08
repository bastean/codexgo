package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *delete.Command {
	id, _ := valueObjectMother.RandomId()

	return &delete.Command{
		Id: id.Value(),
	}
}
