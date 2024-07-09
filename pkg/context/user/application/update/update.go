package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(new *user.User, password *user.Password) error {
	found, err := update.Repository.Search(&model.RepositorySearchCriteria{
		Id: new.Id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(update.Hashing, found.Password.Value, new.Password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if password != nil {
		new.Password = password
	}

	new.Verified = found.Verified

	err = update.Repository.Update(new)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
