package aggregate

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/saggregate"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/domain/message"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type User struct {
	*saggregate.AggregateRoot
	Id       smodel.ValueObject[string]
	Email    smodel.ValueObject[string]
	Username smodel.ValueObject[string]
	Password smodel.ValueObject[string]
	Verified smodel.ValueObject[bool]
}

type UserPrimitive struct {
	Id       string
	Email    string
	Username string
	Password string
	Verified bool
}

func create(id, email, username, password string, verified bool) (*User, error) {
	aggregateRoot := saggregate.NewAggregateRoot()

	idVO, errId := valueobj.NewId(id)
	emailVO, errEmail := valueobj.NewEmail(email)
	usernameVO, errUsername := valueobj.NewUsername(username)
	passwordVO, errPassword := valueobj.NewPassword(password)
	verifiedVO, errVerified := valueobj.NewVerified(verified)

	err := errors.Join(errId, errEmail, errUsername, errPassword, errVerified)

	if err != nil {
		return nil, serror.BubbleUp(err, "create")
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

func FromPrimitives(userPrimitive *UserPrimitive) (*User, error) {
	user, err := create(
		userPrimitive.Id,
		userPrimitive.Email,
		userPrimitive.Username,
		userPrimitive.Password,
		userPrimitive.Verified,
	)

	if err != nil {
		return nil, serror.BubbleUp(err, "FromPrimitives")
	}

	return user, nil
}

func NewUser(id, email, username, password string) (*User, error) {
	verified := false

	user, err := create(
		id,
		email,
		username,
		password,
		verified,
	)

	if err != nil {
		return nil, serror.BubbleUp(err, "NewUser")
	}

	eventMessage, err := message.NewRegisteredSucceededEvent(&message.RegisteredSucceededEventAttributes{
		Id:       id,
		Email:    email,
		Username: username,
	})

	if err != nil {
		return nil, serror.BubbleUp(err, "NewUser")
	}

	user.RecordMessage(eventMessage)

	return user, nil
}
