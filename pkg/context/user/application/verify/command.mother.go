package verify

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.IdWithValidValue()

	return &Command{
		Id: id.Value(),
	}
}
