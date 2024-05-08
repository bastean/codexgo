package commandMother

import (
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
)

func Random() *verify.Command {
	id, _ := valueObjectMother.RandomId()

	return &verify.Command{
		Id: id.Value(),
	}
}
