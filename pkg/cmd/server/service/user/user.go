package user

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/broker"
	"github.com/bastean/codexgo/pkg/cmd/server/service/database"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
)

var userBcryptHashing = cryptographic.NewUserBcryptHashing()
var userMongoRepository = persistence.NewUserMongoRepository(database.Database, userBcryptHashing)

var userRegister = register.NewRegister(userMongoRepository)
var UserRegisterHandler = register.NewCommandHandler(userRegister, broker.Broker)

var userVerify = verify.NewVerify(userMongoRepository)
var UserVerifyHandler = verify.NewCommandHandler(userVerify)

var userLogin = login.NewLogin(userMongoRepository, userBcryptHashing)
var UserLoginHandler = login.NewQueryHandler(userLogin)

var userUpdate = update.NewUpdate(userMongoRepository, userBcryptHashing)
var UserUpdateHandler = update.NewCommandHandler(userUpdate)

var userDelete = delete.NewDelete(userMongoRepository, userBcryptHashing)
var UserDeleteHandler = delete.NewCommandHandler(userDelete)
