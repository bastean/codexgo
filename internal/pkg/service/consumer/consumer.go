package consumer

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

var Service = &struct {
	Notification string
}{
	Notification: log.Service("Notification"),
}

var (
	err error
)

func Start() error {
	log.Starting(Service.Notification)

	if err = InitNotification(); err != nil {
		log.CannotBeStarted(Service.Notification)
		return errors.BubbleUp(err, "Start")
	}

	log.Started(Service.Notification)

	return nil
}
