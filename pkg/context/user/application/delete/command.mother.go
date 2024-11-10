package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func CommandRandomAttributes() *CommandAttributes {
	return &CommandAttributes{
		Id:       user.IdWithValidValue().Value,
		Password: user.PasswordWithValidValue().Value,
	}
}
