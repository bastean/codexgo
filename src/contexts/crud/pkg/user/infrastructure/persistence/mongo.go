package persistence

import (
	"context"

	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/persistence"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "users"

type userDocument struct {
	Id       string `bson:"id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type UserCollection struct {
	collection *mongo.Collection
	hashing    models.Hashing
}

func (db UserCollection) Save(user *aggregate.User) {
	newUser := userDocument(*user.ToPrimitives())

	newUser.Password = db.hashing.Hash(newUser.Password)

	_, err := db.collection.InsertOne(context.Background(), newUser)

	if mongo.IsDuplicateKeyError(err) {
		persistence.HandleDuplicateKeyError(err)
	}
}

func (db UserCollection) Update(user *aggregate.User) {
	updateFilter := bson.M{"id": user.Id.Value}

	updateUser := bson.M{}

	if user.Email != nil {
		updateUser["email"] = user.Email.Value
	}

	if user.Username != nil {
		updateUser["username"] = user.Username.Value
	}

	if user.Password != nil {
		updateUser["password"] = db.hashing.Hash(user.Password.Value)
	}

	_, err := db.collection.UpdateOne(context.Background(), updateFilter, bson.M{"$set": updateUser})

	if err != nil {
		panic(err)
	}
}

func (db UserCollection) Delete(id *sharedVO.Id) {
	deleteFilter := bson.M{"id": id.Value}

	_, err := db.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		panic(err)
	}
}

func (db UserCollection) Search(filter repository.Filter) *aggregate.User {
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

	result := db.collection.FindOne(context.Background(), searchFilter)

	if err := result.Err(); err != nil {
		persistence.HandleDocumentNotFound(index)
	}

	var userPrimitive aggregate.UserPrimitive

	result.Decode(&userPrimitive)

	user := aggregate.FromPrimitives(&userPrimitive)

	return user
}

func NewUserCollection(database *mongo.Database, hashing models.Hashing) *UserCollection {
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

	return &UserCollection{collection, hashing}
}
