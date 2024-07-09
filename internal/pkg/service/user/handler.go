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
	return &create.Handler{
		Create: &create.Create{
			Repository: repository,
		},
		Broker: broker,
	}
}

func NewRead(repository model.Repository) *read.Handler {
	return &read.Handler{
		Read: &read.Read{
			Repository: repository,
		},
	}
}

func NewUpdate(repository model.Repository, hashing model.Hashing) *update.Handler {
	return &update.Handler{
		Update: &update.Update{
			Repository: repository,
			Hashing:    hashing,
		},
	}
}

func NewDelete(repository model.Repository, hashing model.Hashing) *delete.Handler {
	return &delete.Handler{
		Delete: &delete.Delete{
			Repository: repository,
			Hashing:    hashing,
		},
	}
}

func NewVerify(repository model.Repository) *verify.Handler {
	return &verify.Handler{
		Verify: &verify.Verify{
			Repository: repository,
		},
	}
}

func NewLogin(repository model.Repository, hashing model.Hashing) *login.Handler {
	return &login.Handler{
		Login: &login.Login{
			Repository: repository,
			Hashing:    hashing,
		},
	}
}
