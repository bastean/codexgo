package container

import (
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/bastean/codexgo/context/pkg/user/application/register"
	"github.com/bastean/codexgo/context/pkg/user/application/update"
	"github.com/bastean/codexgo/context/pkg/user/infrastructure/persistence"
)

var repository = persistence.Mongo{}

var userRegister = register.NewRegister(repository)
var UserRegisterHandler = register.NewCommandHandler(*userRegister)

var userLogin = login.NewLogin(repository)
var UserLoginHandler = login.NewQueryHandler(*userLogin)

var userUpdate = update.NewUpdate(repository)
var UserUpdateHandler = update.NewCommandHandler(*userUpdate)

var userDelete = delete.NewDelete(repository)
var UserDeleteHandler = delete.NewCommandHandler(*userDelete)
