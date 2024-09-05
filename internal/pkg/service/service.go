package service

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/communication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport"
)

var (
	err error
)

func Up() error {
	if err = env.Init(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	if err = transport.Up(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	if err = communication.Up(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	if err = persistence.Up(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	if err = module.Up(); err != nil {
		return errors.BubbleUp(err, "Up")
	}

	return nil
}

func Down(ctx context.Context) error {
	if err = communication.Down(); err != nil {
		return errors.BubbleUp(err, "Down")
	}

	if err = persistence.Down(ctx); err != nil {
		return errors.BubbleUp(err, "Down")
	}

	return nil
}
