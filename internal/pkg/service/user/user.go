package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

func Start(repository repository.User, broker messages.Broker, hashing hashing.Hashing) {
	Create = &create.Handler{
		Create: &create.Create{
			User: repository,
		},
		Broker: broker,
	}

	Read = &read.Handler{
		Read: &read.Read{
			User: repository,
		},
	}

	Update = &update.Handler{
		Update: &update.Update{
			User:    repository,
			Hashing: hashing,
		},
	}

	Delete = &delete.Handler{
		Delete: &delete.Delete{
			User:    repository,
			Hashing: hashing,
		},
	}

	Verify = &verify.Handler{
		Verify: &verify.Verify{
			User: repository,
		},
	}

	Login = &login.Handler{
		Login: &login.Login{
			User:    repository,
			Hashing: hashing,
		},
	}
}
