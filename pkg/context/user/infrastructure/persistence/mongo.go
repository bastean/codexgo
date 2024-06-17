package persistence

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDocument struct {
	Id       string `bson:"id,omitempty"`
	Email    string `bson:"email,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Verified bool   `bson:"verified,omitempty"`
}

type UserCollection struct {
	collection *mongo.Collection
	hashing    model.Hashing
}

func (db *UserCollection) Save(user *aggregate.User) error {
	newUser := UserDocument(*user.ToPrimitives())

	hashed, err := db.hashing.Hash(newUser.Password)

	if err != nil {
		return errors.BubbleUp(err, "Save")
	}

	newUser.Password = hashed

	_, err = db.collection.InsertOne(context.Background(), &newUser)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(persistences.HandleMongoDuplicateKeyError(err), "Save")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a user",
			Why: errors.Meta{
				"Id": user.Id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Verify(id models.ValueObject[string]) error {
	filter := bson.D{{Key: "id", Value: id.Value()}}

	_, err := db.collection.UpdateOne(context.Background(), filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "verified", Value: true},
		}},
	})

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Verify",
			What:  "failure to verify a user",
			Why: errors.Meta{
				"Id": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Update(user *aggregate.User) error {
	updatedUser := UserDocument(*user.ToPrimitives())

	filter := bson.D{{Key: "id", Value: user.Id.Value()}}

	hashed, err := db.hashing.Hash(user.Password.Value())

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	updatedUser.Password = hashed

	_, err = db.collection.ReplaceOne(context.Background(), filter, &updatedUser)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a user",
			Why: errors.Meta{
				"Id": user.Id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Delete(id models.ValueObject[string]) error {
	filter := bson.D{{Key: "id", Value: id.Value()}}

	_, err := db.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a user",
			Why: errors.Meta{
				"Id": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Search(criteria *model.RepositorySearchCriteria) (*aggregate.User, error) {
	var filter bson.D
	var index string

	switch {
	case criteria.Id != nil:
		filter = bson.D{{Key: "id", Value: criteria.Id.Value()}}
		index = criteria.Id.Value()
	case criteria.Email != nil:
		filter = bson.D{{Key: "email", Value: criteria.Email.Value()}}
		index = criteria.Email.Value()
	}

	result := db.collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, persistences.HandleMongoDocumentNotFound(index, err)
	}

	primitive := new(aggregate.UserPrimitive)

	err := result.Decode(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to decode a result",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	user, err := aggregate.FromPrimitives(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to create an aggregate from a primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Index":     index,
			},
			Who: err,
		})
	}

	return user, nil
}

func NewMongoCollection(mdb *persistences.MongoDB, collectionName string, hashing model.Hashing) (model.Repository, error) {
	collection := mdb.Database.Collection(collectionName)

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
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

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMongoCollection",
			What:  "failure to create indexes for user collection",
			Why: errors.Meta{
				"Collection": collectionName,
			},
			Who: err,
		})
	}

	return &UserCollection{
		collection: collection,
		hashing:    hashing,
	}, nil
}
