package user

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/broker"
	"github.com/bastean/codexgo/pkg/cmd/server/service/database"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
)

var Bcrypt = new(cryptographic.Bcrypt)

var CollectionName = "users"
var MongoCollection model.Repository

var Register *register.Register
var RegisterHandler *register.CommandHandler

var Verify *verify.Verify
var VerifyHandler *verify.CommandHandler

var Login *login.Login
var LoginHandler *login.QueryHandler

var Update *update.Update
var UpdateHandler *update.CommandHandler

var Delete *delete.Delete
var DeleteHandler *delete.CommandHandler

func Init() error {
	collection, err := persistence.NewMongoCollection(database.Database, CollectionName, Bcrypt)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	MongoCollection = collection

	Register = &register.Register{
		Repository: MongoCollection,
	}
	RegisterHandler = &register.CommandHandler{
		UseCase: Register,
		Broker:  broker.Broker,
	}

	Verify = &verify.Verify{
		Repository: MongoCollection,
	}
	VerifyHandler = &verify.CommandHandler{
		UseCase: Verify,
	}

	Login = &login.Login{
		Repository: MongoCollection,
		Hashing:    Bcrypt,
	}
	LoginHandler = &login.QueryHandler{
		UseCase: Login,
	}

	Update = &update.Update{
		Repository: MongoCollection,
		Hashing:    Bcrypt,
	}
	UpdateHandler = &update.CommandHandler{
		UseCase: Update,
	}

	Delete = &delete.Delete{
		Repository: MongoCollection,
		Hashing:    Bcrypt,
	}
	DeleteHandler = &delete.CommandHandler{
		UseCase: Delete,
	}

	return nil
}
