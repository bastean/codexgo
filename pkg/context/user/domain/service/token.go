package service

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func SetResetToken(aggregate *user.User, id string) error {
	if !aggregate.HasResetToken() {
		resetToken, err := token.New(id)

		if err != nil {
			return errors.BubbleUp(err)
		}

		aggregate.ResetToken = resetToken

		return nil
	}

	resetTokenID, err := values.Replace(aggregate.ResetToken.ID, id)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = aggregate.ResetToken.Attempt.Increase()

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.ResetToken.ID = resetTokenID

	return nil
}
