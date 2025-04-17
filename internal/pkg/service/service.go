package service

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/authentication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/consumer"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/handler"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

var (
	err error
)

func Up() error {
	if err = env.Init(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = authentication.Up(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = transport.Up(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = communication.Up(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = persistence.Up(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = consumer.Start(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = handler.Start(); err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}

func Down(ctx context.Context) error {
	if err = communication.Down(); err != nil {
		return errors.BubbleUp(err)
	}

	if err = persistence.Down(ctx); err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
