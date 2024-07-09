package read

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

func RandomQuery() *Query {
	id := user.IdWithValidValue()

	return &Query{
		Id: id.Value,
	}
}
