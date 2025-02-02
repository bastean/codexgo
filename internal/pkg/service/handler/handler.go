package handler

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

var Service = &struct {
	User string
}{
	User: log.Service("User"),
}

var (
	err error
)

func Start() error {
	log.Starting(Service.User)

	if err = InitUser(); err != nil {
		log.CannotBeStarted(Service.User)
		return errors.BubbleUp(err, "Start")
	}

	log.Started(Service.User)

	return nil
}
