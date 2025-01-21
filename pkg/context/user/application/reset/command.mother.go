package reset

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func CommandRandomAttributes() *CommandAttributes {
	return &CommandAttributes{
		Reset:    user.IDWithValidValue().Value,
		ID:       user.IDWithValidValue().Value,
		Password: user.PlainPasswordWithValidValue().Value,
	}
}
