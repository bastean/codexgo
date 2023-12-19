package persistence

import (
	"context"

	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
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
}

func (mongo Mongo) Save(user *aggregate.User) {
	newUser := userDocument(*user.ToPrimitives())

	_, err := mongo.collection.InsertOne(context.Background(), newUser)

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Update(user *aggregate.User) {
	updateFilter := bson.M{"id": user.Id.Value}
	userPrimitives := *user.ToPrimitives()
	updateUser := bson.M{"$set": bson.M{
		"email":    userPrimitives.Email,
		"username": userPrimitives.Username,
		"password": userPrimitives.Password,
	}}

	_, err := mongo.collection.UpdateOne(context.Background(), updateFilter, updateUser)

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Delete(email *sharedVO.Email) {
	deleteFilter := bson.M{"email": email.Value}

	_, err := mongo.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		panic(err)
	}
}

func (mongo Mongo) Search(email *sharedVO.Email) (*aggregate.User, error) {
	searchFilter := bson.M{"email": email.Value}

	result := mongo.collection.FindOne(context.Background(), searchFilter)

	var userPrimitive aggregate.UserPrimitive

	result.Decode(&userPrimitive)

	user, _ := aggregate.FromPrimitives(&userPrimitive)

	return user, nil
}

func NewMongo() *Mongo {
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

	return &Mongo{collection: collection}
}
