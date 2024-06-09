package read

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	id, _ := valueobj.IdWithValidValue()

	return &Query{
		Id: id.Value(),
	}
}
