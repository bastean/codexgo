package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type (
	CreateCommand = create.Command
	UpdateCommand = update.Command
	DeleteCommand = delete.Command
	VerifyCommand = verify.Command
)

type (
	ReadQuery  = read.Query
	LoginQuery = login.Query
)

type (
	ReadResponse = read.Response
)

func NewCreate(repository model.Repository, broker messages.Broker) *create.Handler {
	usecase := &create.Create{
		Repository: repository,
	}

	return &create.Handler{
		UseCase: usecase,
		Broker:  broker,
	}
}

func NewRead(repository model.Repository) *read.Handler {
	usecase := &read.Read{
		Repository: repository,
	}

	return &read.Handler{
		UseCase: usecase,
	}
}

func NewUpdate(repository model.Repository, hashing model.Hashing) *update.Handler {
	usecase := &update.Update{
		Repository: repository,
		Hashing:    hashing,
	}

	return &update.Handler{
		UseCase: usecase,
	}
}

func NewDelete(repository model.Repository, hashing model.Hashing) *delete.Handler {
	usecase := &delete.Delete{
		Repository: repository,
		Hashing:    hashing,
	}

	return &delete.Handler{
		UseCase: usecase,
	}
}

func NewVerify(repository model.Repository) *verify.Handler {
	usecase := &verify.Verify{
		Repository: repository,
	}

	return &verify.Handler{
		UseCase: usecase,
	}
}

func NewLogin(repository model.Repository, hashing model.Hashing) *login.Handler {
	usecase := &login.Login{
		Repository: repository,
		Hashing:    hashing,
	}

	return &login.Handler{
		UseCase: usecase,
	}
}
