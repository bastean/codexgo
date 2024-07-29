package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func RandomCommand() *Command {
	id := user.IdWithValidValue()

	return &Command{
		Id: id.Value,
	}
}
