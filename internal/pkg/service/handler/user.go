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
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/badgerdb/kv"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
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
		repository, err = kv.Open(
			persistence.BadgerDB,
		)
	}

	if err != nil {
		return errors.BubbleUp(err)
	}

	UserCreate = &create.Handler{
		Case: &create.Case{
			Hasher:     cipher.Hasher,
			Repository: repository,
			EventBus:   event.Bus,
		},
	}

	UserUpdate = &update.Handler{
		Case: &update.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserDelete = &delete.Handler{
		Case: &delete.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserVerify = &verify.Handler{
		Case: &verify.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserForgot = &forgot.Handler{
		Case: &forgot.Case{
			Repository: repository,
			EventBus:   event.Bus,
		},
	}

	UserReset = &reset.Handler{
		Case: &reset.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	UserRead = &read.Handler{
		Case: &read.Case{
			Repository: repository,
		},
	}

	UserLogin = &login.Handler{
		Case: &login.Case{
			Repository: repository,
			Hasher:     cipher.Hasher,
		},
	}

	err = commands.AddCommandMapper(command.Bus, commands.Mapper{
		create.CommandKey.Value(): UserCreate,
		update.CommandKey.Value(): UserUpdate,
		delete.CommandKey.Value(): UserDelete,
		verify.CommandKey.Value(): UserVerify,
		forgot.CommandKey.Value(): UserForgot,
		reset.CommandKey.Value():  UserReset,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = queries.AddQueryMapper(query.Bus, queries.Mapper{
		read.QueryKey.Value():  UserRead,
		login.QueryKey.Value(): UserLogin,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
