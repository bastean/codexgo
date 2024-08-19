package collection

import (
	"context"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	*mongo.Collection
	hashing.Hashing
}

func (mongoDB *User) Create(user *user.User) error {
	new := user.ToPrimitive()

	hashed, err := mongoDB.Hashing.Hash(new.Password)

	if err != nil {
		return errors.BubbleUp(err, "Create")
	}

	new.Password = hashed

	_, err = mongoDB.Collection.InsertOne(context.Background(), new)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(mongodb.HandleDuplicateKeyError(err), "Create")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Create",
			What:  "Failure to create a User",
			Why: errors.Meta{
				"Id": user.Id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Verify(id *user.Id) error {
	filter := bson.D{{Key: "id", Value: id.Value}}

	_, err := mongoDB.Collection.UpdateOne(context.Background(), filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "verified", Value: true},
		}},
	})

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Verify",
			What:  "Failure to verify a User",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Update(user *user.User) error {
	updated := user.ToPrimitive()

	filter := bson.D{{Key: "id", Value: user.Id.Value}}

	hashed, err := mongoDB.Hashing.Hash(user.Password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	updated.Password = hashed

	_, err = mongoDB.Collection.ReplaceOne(context.Background(), filter, updated)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "Failure to update a User",
			Why: errors.Meta{
				"Id": user.Id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Delete(id *user.Id) error {
	filter := bson.D{{Key: "id", Value: id.Value}}

	_, err := mongoDB.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "Failure to delete a User",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Search(criteria *repository.SearchCriteria) (*user.User, error) {
	var filter bson.D
	var index string

	switch {
	case criteria.Id != nil:
		filter = bson.D{{Key: "id", Value: criteria.Id.Value}}
		index = criteria.Id.Value
	case criteria.Email != nil:
		filter = bson.D{{Key: "email", Value: criteria.Email.Value}}
		index = criteria.Email.Value
	}

	result := mongoDB.Collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, mongodb.HandleDocumentNotFound(index, err)
	}

	primitive := new(user.Primitive)

	err := result.Decode(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to decode a result",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	found, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to create a User from a Primitive",
			Why: errors.Meta{
				"Index":     index,
				"Primitive": primitive,
			},
			Who: err,
		})
	}

	return found, nil
}

func OpenUser(session *mongodb.MongoDB, name string, hashing hashing.Hashing) (repository.Repository, error) {
	collection := session.Database.Collection(name)

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
			Where: "OpenUser",
			What:  "Failure to create Indexes for User Collection",
			Why: errors.Meta{
				"Collection": name,
			},
			Who: err,
		})
	}

	return &User{
		Collection: collection,
		Hashing:    hashing,
	}, nil
}
