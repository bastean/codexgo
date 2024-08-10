package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/service"
)

type Update struct {
	repository.Repository
	hashing.Hashing
}

func (update *Update) Run(account *user.User, updated *user.Password) error {
	found, err := update.Repository.Search(&repository.SearchCriteria{
		Id: account.Id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(update.Hashing, found.Password.Value, account.Password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if updated != nil {
		account.Password = updated
	}

	account.Verified = found.Verified

	err = update.Repository.Update(account)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
