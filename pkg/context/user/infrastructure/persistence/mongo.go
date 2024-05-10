package persistence

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/spersistence"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userDocument struct {
	Id       string `bson:"id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Verified bool   `bson:"verified"`
}

type UserCollection struct {
	collection *mongo.Collection
	hashing    model.Hashing
}

func (db *UserCollection) Save(user *aggregate.User) error {
	newUser := userDocument(*user.ToPrimitives())

	hashed, err := db.hashing.Hash(newUser.Password)

	if err != nil {
		return serror.BubbleUp(err, "Save")
	}

	newUser.Password = hashed

	_, err = db.collection.InsertOne(context.Background(), newUser)

	if mongo.IsDuplicateKeyError(err) {
		return serror.BubbleUp(spersistence.HandleMongoDuplicateKeyError(err), "Save")
	}

	return nil
}

func (db *UserCollection) Update(user *aggregate.User) error {
	updateFilter := bson.M{"id": user.Id.Value()}

	updateUser := bson.M{}

	if user.Email != nil {
		updateUser["email"] = user.Email.Value()
	}

	if user.Username != nil {
		updateUser["username"] = user.Username.Value()
	}

	if user.Password != nil {
		hashed, err := db.hashing.Hash(user.Password.Value())

		if err != nil {
			return serror.BubbleUp(err, "Update")
		}

		updateUser["password"] = hashed
	}

	if user.Verified != nil {
		updateUser["verified"] = user.Verified.Value()
	}

	_, err := db.collection.UpdateOne(context.Background(), updateFilter, bson.M{"$set": updateUser})

	if err != nil {
		return serror.NewFailedError(&serror.Bubble{
			Where: "Update",
			What:  "failed on update",
			Why: serror.Meta{
				"Id": user.Id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Delete(id smodel.ValueObject[string]) error {
	deleteFilter := bson.M{"id": id.Value()}

	_, err := db.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		return serror.NewFailedError(&serror.Bubble{
			Where: "Delete",
			What:  "failed on delete",
			Why: serror.Meta{
				"Id": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Search(filter model.RepositorySearchCriteria) (*aggregate.User, error) {
	var searchFilter bson.M
	var index string

	if filter.Email != nil {
		searchFilter = bson.M{"email": filter.Email.Value()}
		index = filter.Email.Value()
	}

	if filter.Id != nil {
		searchFilter = bson.M{"id": filter.Id.Value()}
		index = filter.Id.Value()
	}

	result := db.collection.FindOne(context.Background(), searchFilter)

	if err := result.Err(); err != nil {
		spersistence.HandleMongoDocumentNotFound(index, err)
	}

	var userPrimitive aggregate.UserPrimitive

	result.Decode(&userPrimitive)

	user, err := aggregate.FromPrimitives(&userPrimitive)

	if err != nil {
		return nil, serror.NewFailedError(&serror.Bubble{
			Where: "Search",
			What:  "failed on search",
			Why: serror.Meta{
				"Id":    filter.Id.Value(),
				"Email": filter.Email.Value(),
			},
			Who: err,
		})
	}

	return user, nil
}

func NewUserMongoRepository(mdb *spersistence.MongoDB, collectionName string, hashing model.Hashing) model.Repository {
	collection := mdb.Database.Collection(collectionName)

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

	return &UserCollection{
		collection: collection,
		hashing:    hashing,
	}
}
