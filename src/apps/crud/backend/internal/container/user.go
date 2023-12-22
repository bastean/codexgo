package container

import (
	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/bastean/codexgo/context/pkg/user/application/register"
	"github.com/bastean/codexgo/context/pkg/user/application/update"
	"github.com/bastean/codexgo/context/pkg/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/context/pkg/user/infrastructure/persistence"
)

var hashing = cryptographic.Bcrypt{}
var repository = persistence.NewMongo(hashing)

var userRegister = register.NewRegister(repository)
var UserRegisterHandler = register.NewCommandHandler(*userRegister)

var userLogin = login.NewLogin(repository, hashing)
var UserLoginHandler = login.NewQueryHandler(*userLogin)

var userUpdate = update.NewUpdate(repository, hashing)
var UserUpdateHandler = update.NewCommandHandler(*userUpdate)

var userDelete = delete.NewDelete(repository, hashing)
var UserDeleteHandler = delete.NewCommandHandler(*userDelete)
