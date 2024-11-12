package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func QueryRandomAttributes() *QueryAttributes {
	return &QueryAttributes{
		ID: user.IDWithValidValue().Value,
	}
}
