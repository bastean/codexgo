package user

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/broker"
	"github.com/bastean/codexgo/pkg/cmd/server/service/database"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
)

var Bcrypt = new(cryptographic.Bcrypt)

var CollectionName = "users"
var MongoCollection model.Repository

var Create *create.Create
var CreateHandler *create.CommandHandler

var Read *read.Read
var ReadHandler *read.QueryHandler

var Update *update.Update
var UpdateHandler *update.CommandHandler

var Delete *delete.Delete
var DeleteHandler *delete.CommandHandler

var Verify *verify.Verify
var VerifyHandler *verify.CommandHandler

var Login *login.Login
var LoginHandler *login.QueryHandler

func Init() error {
	logger.Info("starting module: user")

	collection, err := persistence.NewMongoCollection(database.Database, CollectionName, Bcrypt)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	MongoCollection = collection

	Create = &create.Create{
		Repository: MongoCollection,
	}
	CreateHandler = &create.CommandHandler{
		UseCase: Create,
		Broker:  broker.Broker,
	}

	Read = &read.Read{
		Repository: MongoCollection,
	}
	ReadHandler = &read.QueryHandler{
		UseCase: Read,
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

	return nil
}
