package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObjects"
)

func Create(id *sharedVO.Id) *delete.Command {
	return &delete.Command{Id: id.Value}
}

func Random() *delete.Command {
	return Create(create.RandomId())
}

func Invalid() *delete.Command {
	return Create(create.InvalidId())
}