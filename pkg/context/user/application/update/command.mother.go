package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func CommandRandomAttributes() *CommandAttributes {
	return &CommandAttributes{
		Id:              user.IdWithValidValue().Value,
		Email:           user.EmailWithValidValue().Value,
		Username:        user.UsernameWithValidValue().Value,
		Password:        user.PasswordWithValidValue().Value,
		UpdatedPassword: user.PasswordWithValidValue().Value,
	}
}
