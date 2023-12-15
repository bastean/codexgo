package container

import (
	"github.com/bastean/codexgo/context/pkg/user/application/register"
	"github.com/bastean/codexgo/context/pkg/user/infrastructure/persistence"
)

var repository = persistence.Mongo{}

var userRegister = register.NewRegister(repository)

var UserRegisterHandler = register.NewCommandHandler(*userRegister)
