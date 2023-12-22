package persistence

import (
	"context"

	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://codexgo:codexgo@localhost:27017" //! .env
var databaseName = "codexgo"                          //! .env
var collectionName = "users"                          //! .env

type userDocument struct {
	Id       string `bson:"id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Mongo struct {
	collection *mongo.Collection
	hashing    models.Hashing
}

func (mongo Mongo) Save(user *aggregate.User) {
	newUser := userDocument(*user.ToPrimitives())

	newUser.Password = mongo.hashing.Hash(newUser.Password)

	_, err := mongo.collection.InsertOne(context.Background(), newUser)

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Update(user *aggregate.User) {
	updateFilter := bson.M{"id": user.Id.Value}

	updateUser := bson.M{}

	if user.Email != nil {
		updateUser["email"] = user.Email.Value
	}

	if user.Username != nil {
		updateUser["username"] = user.Username.Value
	}

	if user.Password != nil {
		updateUser["password"] = mongo.hashing.Hash(user.Password.Value)
	}

	_, err := mongo.collection.UpdateOne(context.Background(), updateFilter, bson.M{"$set": updateUser})

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Delete(id *sharedVO.Id) {
	deleteFilter := bson.M{"id": id.Value}

	_, err := mongo.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Search(filter repository.Filter) *aggregate.User {
	var searchFilter bson.M

	if filter.Email != nil {
		searchFilter = bson.M{"email": filter.Email.Value}
	}

	if filter.Id != nil {
		searchFilter = bson.M{"id": filter.Id.Value}
	}

	result := mongo.collection.FindOne(context.Background(), searchFilter)

	if err := result.Err(); err != nil {
		panic("not found")
	}

	var userPrimitive aggregate.UserPrimitive

	result.Decode(&userPrimitive)

	user, _ := aggregate.FromPrimitives(&userPrimitive)

	return user
}

func NewMongo(hashing models.Hashing) *Mongo {
	var err error

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	database := client.Database(databaseName)
	collection := database.Collection(collectionName)

	return &Mongo{collection: collection, hashing: hashing}
}
