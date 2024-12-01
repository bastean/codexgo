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

func Start(repository repository.Repository, bus event.Bus, hasher hashes.Hasher) {
	CreateHandler = &create.Handler{
		Create: &create.Case{
			Hasher:     hasher,
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
			Hasher:     hasher,
		},
	}

	DeleteHandler = &delete.Handler{
		Delete: &delete.Case{
			Repository: repository,
			Hasher:     hasher,
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
			Hasher:     hasher,
		},
	}
}
