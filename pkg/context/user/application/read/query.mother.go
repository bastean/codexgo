package read

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	id, _ := valueobj.RandomId()

	return &Query{
		Id: id.Value(),
	}
}
