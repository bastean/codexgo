package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Update struct {
	repository.User
	hashing.Hashing
}

func (update *Update) Run(new *user.User, password *user.Password) error {
	found, err := update.User.Search(&repository.UserSearchCriteria{
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

	err = update.User.Update(new)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
