package service

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func UpdateEmail(email string, aggregate *user.User) error {
	if email != "" && email != aggregate.Email.Value() {
		_, err := values.New[*user.Email](email)

		if err != nil {
			return errors.BubbleUp(err)
		}

		aggregate.Email, err = values.Replace(aggregate.Email, email)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	return nil
}

func UpdateUsername(username string, aggregate *user.User) error {
	if username != "" && username != aggregate.Username.Value() {
		_, err := values.New[*user.Username](username)

		if err != nil {
			return errors.BubbleUp(err)
		}

		aggregate.Username, err = values.Replace(aggregate.Username, username)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	return nil
}

func UpdatePassword(password string, aggregate *user.User, hasher roles.Hasher) error {
	if password != "" {
		_, err := values.New[*user.PlainPassword](password)

		if err != nil {
			return errors.BubbleUp(err)
		}

		password, err = hasher.Hash(password)

		if err != nil {
			return errors.BubbleUp(err)
		}

		aggregate.Password, err = values.Replace(aggregate.Password, password)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	return nil
}
