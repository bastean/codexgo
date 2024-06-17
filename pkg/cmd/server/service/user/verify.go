package user

import (
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type VerifyCommand = verify.Command

func NewVerify(repository model.Repository) *verify.Handler {
	usecase := &verify.Verify{
		Repository: repository,
	}

	return &verify.Handler{
		UseCase: usecase,
	}
}
