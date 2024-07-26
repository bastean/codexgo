package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
)

var (
	Create *create.Handler
	Read   *read.Handler
	Update *update.Handler
	Delete *delete.Handler
	Verify *verify.Handler
	Login  *login.Handler
)

func Start(repository repository.User, broker messages.Broker, hashing hashing.Hashing) {
	Create = NewCreate(repository, broker)

	Read = NewRead(repository)

	Update = NewUpdate(repository, hashing)

	Delete = NewDelete(repository, hashing)

	Verify = NewVerify(repository)

	Login = NewLogin(repository, hashing)
}
