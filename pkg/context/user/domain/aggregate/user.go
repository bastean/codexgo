package aggregate

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type User struct {
	*aggregates.AggregateRoot
	Id, Email, Username, Password models.ValueObject[string]
	Verified                      models.ValueObject[bool]
}

type UserPrimitive struct {
	Id, Email, Username, Password string
	Verified                      bool
}

func create(primitive *UserPrimitive) (*User, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	idVO, errId := valueobj.NewId(primitive.Id)
	emailVO, errEmail := valueobj.NewEmail(primitive.Email)
	usernameVO, errUsername := valueobj.NewUsername(primitive.Username)
	passwordVO, errPassword := valueobj.NewPassword(primitive.Password)
	verifiedVO, errVerified := valueobj.NewVerified(primitive.Verified)

	err := errors.Join(errId, errEmail, errUsername, errPassword, errVerified)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		AggregateRoot: aggregateRoot,
		Id:            idVO,
		Email:         emailVO,
		Username:      usernameVO,
		Password:      passwordVO,
		Verified:      verifiedVO,
	}, nil
}

func (user *User) ToPrimitives() *UserPrimitive {
	return &UserPrimitive{
		Id:       user.Id.Value(),
		Email:    user.Email.Value(),
		Username: user.Username.Value(),
		Password: user.Password.Value(),
		Verified: user.Verified.Value(),
	}
}

func FromPrimitives(primitive *UserPrimitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return user, nil
}

func NewUser(primitive *UserPrimitive) (*User, error) {
	primitive.Verified = false

	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewUser")
	}

	message, err := event.NewCreatedSucceeded(&event.CreatedSucceeded{
		Attributes: &event.CreatedSucceededAttributes{
			Id:       user.Id.Value(),
			Email:    user.Email.Value(),
			Username: user.Username.Value(),
		},
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewUser")
	}

	user.RecordMessage(message)

	return user, nil
}
