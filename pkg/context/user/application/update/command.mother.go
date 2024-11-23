package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func CommandRandomAttributes() *CommandAttributes {
	return &CommandAttributes{
		ID:              user.IDWithValidValue().Value,
		Email:           user.EmailWithValidValue().Value,
		Username:        user.UsernameWithValidValue().Value,
		Password:        user.PlainPasswordWithValidValue().Value,
		UpdatedPassword: user.PlainPasswordWithValidValue().Value,
	}
}
