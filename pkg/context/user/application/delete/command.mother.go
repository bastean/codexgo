package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func CommandRandomAttributes() *CommandAttributes {
	return &CommandAttributes{
		ID:       user.IDWithValidValue().Value,
		Password: user.PlainPasswordWithValidValue().Value,
	}
}
