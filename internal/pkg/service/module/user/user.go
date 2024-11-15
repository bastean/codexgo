package user

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/event"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

func Start(repository repository.Repository, bus event.Bus, hashing hashes.Hashing) {
	CreateHandler = &create.Handler{
		Create: &create.Case{
			Repository: repository,
		},
		Bus: bus,
	}

	ReadHandler = &read.Handler{
		Read: &read.Case{
			Repository: repository,
		},
	}

	UpdateHandler = &update.Handler{
		Update: &update.Case{
			Repository: repository,
			Hashing:    hashing,
		},
	}

	DeleteHandler = &delete.Handler{
		Delete: &delete.Case{
			Repository: repository,
			Hashing:    hashing,
		},
	}

	VerifyHandler = &verify.Handler{
		Verify: &verify.Case{
			Repository: repository,
		},
	}

	LoginHandler = &login.Handler{
		Login: &login.Case{
			Repository: repository,
			Hashing:    hashing,
		},
	}
}
