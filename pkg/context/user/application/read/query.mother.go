package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func RandomQuery() *Query {
	id := user.IdWithValidValue()

	return &Query{
		Id: id.Value,
	}
}
