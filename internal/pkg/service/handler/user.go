package handler

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/cipher"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/event"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/forgot"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/reset"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/sqlite/table"
)

const (
	UserCollectionName = "users"
)

var (
	UserCreate *create.Handler
	UserUpdate *update.Handler
	UserDelete *delete.Handler
	UserVerify *verify.Handler
	UserForgot *forgot.Handler
	UserReset  *reset.Handler
)

var (
	UserRead  *read.Handler
	UserLogin *login.Handler
)

func InitUser() error {
	var (
		repository role.Repository
	)

	switch {
	case persistence.MongoDB != nil:
		repository, err = collection.Open(
			persistence.MongoDB,
			UserCollectionName,
		)
	default:
		repository, err = table.Open(
			persistence.SQLite,
		)
	}

	if err != nil {
		return errors.BubbleUp(err, "InitUser")
	}

	UserCreate = &create.Handler{
		Create: &create.Case{
			Hasher:     cipher.Hasher,
			Repository: repository,
		},
		EventBus: event.Bus,
	}

	UserUpdate = &update.Handler{
		Update: &update.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserDelete = &delete.Handler{
		Delete: &delete.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserVerify = &verify.Handler{
		Verify: &verify.Case{
			Repository: repository,
		},
	}

	UserForgot = &forgot.Handler{
		Forgot: &forgot.Case{
			Repository: repository,
		},
		EventBus: event.Bus,
	}

	UserReset = &reset.Handler{
		Reset: &reset.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserRead = &read.Handler{
		Read: &read.Case{
			Repository: repository,
		},
	}

	UserLogin = &login.Handler{
		Login: &login.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	err = commands.AddCommandMapper(command.Bus, commands.Mapper{
		create.CommandKey: UserCreate,
		update.CommandKey: UserUpdate,
		delete.CommandKey: UserDelete,
		verify.CommandKey: UserVerify,
		forgot.CommandKey: UserForgot,
		reset.CommandKey:  UserReset,
	})

	if err != nil {
		return errors.BubbleUp(err, "InitUser")
	}

	err = queries.AddQueryMapper(query.Bus, queries.Mapper{
		read.QueryKey:  UserRead,
		login.QueryKey: UserLogin,
	})

	if err != nil {
		return errors.BubbleUp(err, "InitUser")
	}

	return nil
}
