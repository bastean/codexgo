package read

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	id := valueobj.IdWithValidValue()

	return &Query{
		Id: id.Value(),
	}
}
