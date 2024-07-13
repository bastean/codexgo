package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
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

func NewCreate(repository repository.User, broker messages.Broker) *create.Handler {
	return &create.Handler{
		Create: &create.Create{
			User: repository,
		},
		Broker: broker,
	}
}

func NewRead(repository repository.User) *read.Handler {
	return &read.Handler{
		Read: &read.Read{
			User: repository,
		},
	}
}

func NewUpdate(repository repository.User, hashing hashing.Hashing) *update.Handler {
	return &update.Handler{
		Update: &update.Update{
			User:    repository,
			Hashing: hashing,
		},
	}
}

func NewDelete(repository repository.User, hashing hashing.Hashing) *delete.Handler {
	return &delete.Handler{
		Delete: &delete.Delete{
			User:    repository,
			Hashing: hashing,
		},
	}
}

func NewVerify(repository repository.User) *verify.Handler {
	return &verify.Handler{
		Verify: &verify.Verify{
			User: repository,
		},
	}
}

func NewLogin(repository repository.User, hashing hashing.Hashing) *login.Handler {
	return &login.Handler{
		Login: &login.Login{
			User:    repository,
			Hashing: hashing,
		},
	}
}
