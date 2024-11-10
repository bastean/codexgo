package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

func Start(repository repository.Repository, bus events.Bus, hashing hashing.Hashing) {
	CreateHandler = &create.Handler{
		Create: &create.Create{
			Repository: repository,
		},
		Bus: bus,
	}

	ReadHandler = &read.Handler{
		Read: &read.Read{
			Repository: repository,
		},
	}

	UpdateHandler = &update.Handler{
		Update: &update.Update{
			Repository: repository,
			Hashing:    hashing,
		},
	}

	DeleteHandler = &delete.Handler{
		Delete: &delete.Delete{
			Repository: repository,
			Hashing:    hashing,
		},
	}

	VerifyHandler = &verify.Handler{
		Verify: &verify.Verify{
			Repository: repository,
		},
	}

	LoginHandler = &login.Handler{
		Login: &login.Login{
			Repository: repository,
			Hashing:    hashing,
		},
	}
}
