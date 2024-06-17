package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"

	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

var Create *create.Handler

var Read *read.Handler

var Update *update.Handler

var Delete *delete.Handler

var Verify *verify.Handler

var Login *login.Handler

func Init(repository model.Repository, broker models.Broker, hashing model.Hashing) {
	Create = NewCreate(repository, broker)

	Read = NewRead(repository)

	Update = NewUpdate(repository, hashing)

	Delete = NewDelete(repository, hashing)

	Verify = NewVerify(repository)

	Login = NewLogin(repository, hashing)
}
