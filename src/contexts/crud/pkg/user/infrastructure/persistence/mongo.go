package persistence

import (
	"context"
	"os"

	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = os.Getenv("DATABASE_URI")
var databaseName = os.Getenv("DATABASE_NAME")
var collectionName = os.Getenv("DATABASE_COLLECTION")

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

func (mg Mongo) Save(user *aggregate.User) {
	newUser := userDocument(*user.ToPrimitives())

	newUser.Password = mg.hashing.Hash(newUser.Password)

	_, err := mg.collection.InsertOne(context.Background(), newUser)

	if mongo.IsDuplicateKeyError(err) {
		handleDuplicateKeyError(err)
	}
}

func (mg Mongo) Update(user *aggregate.User) {
	updateFilter := bson.M{"id": user.Id.Value}

	updateUser := bson.M{}

	if user.Email != nil {
		updateUser["email"] = user.Email.Value
	}

	if user.Username != nil {
		updateUser["username"] = user.Username.Value
	}

	if user.Password != nil {
		updateUser["password"] = mg.hashing.Hash(user.Password.Value)
	}

	_, err := mg.collection.UpdateOne(context.Background(), updateFilter, bson.M{"$set": updateUser})

	if err != nil {
		panic(err)
	}
}

func (mg Mongo) Delete(id *sharedVO.Id) {
	deleteFilter := bson.M{"id": id.Value}

	_, err := mg.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		panic(err)
	}
}

func (mg Mongo) Search(filter repository.Filter) *aggregate.User {
	var searchFilter bson.M
	var index string

	if filter.Email != nil {
		searchFilter = bson.M{"email": filter.Email.Value}
		index = filter.Email.Value
	}

	if filter.Id != nil {
		searchFilter = bson.M{"id": filter.Id.Value}
		index = filter.Id.Value
	}

	result := mg.collection.FindOne(context.Background(), searchFilter)

	if err := result.Err(); err != nil {
		handleDocumentNotFound(index)
	}

	var userPrimitive aggregate.UserPrimitive

	result.Decode(&userPrimitive)

	user := aggregate.FromPrimitives(&userPrimitive)

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

	collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	return &Mongo{collection: collection, hashing: hashing}
}
