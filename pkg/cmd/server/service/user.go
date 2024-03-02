package service

import (
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
)

var bcrypt = cryptographic.Bcrypt{}
var userCollection = persistence.NewUserCollection(Database, bcrypt)

var userRegister = register.NewRegister(userCollection)
var UserRegisterHandler = register.NewCommandHandler(*userRegister)

var userLogin = login.NewLogin(userCollection, bcrypt)
var UserLoginHandler = login.NewQueryHandler(*userLogin)

var userUpdate = update.NewUpdate(userCollection, bcrypt)
var UserUpdateHandler = update.NewCommandHandler(*userUpdate)

var userDelete = delete.NewDelete(userCollection, bcrypt)
var UserDeleteHandler = delete.NewCommandHandler(*userDelete)
