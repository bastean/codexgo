package recipient

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Recipient struct {
	VerifyToken, ResetToken *ID
	*ID
	*Email
	*Username
}

type Required struct {
	ID, Email, Username string
}

func New(required *Required) (*Recipient, error) {
	id, errID := values.New[*ID](required.ID)

	email, errEmail := values.New[*Email](required.Email)
	username, errUsername := values.New[*Username](required.Username)

	if err := errors.Join(errID, errEmail, errUsername); err != nil {
		return nil, errors.BubbleUp(err)
	}

	recipient := &Recipient{
		ID:       id,
		Email:    email,
		Username: username,
	}

	return recipient, nil
}
