package collection

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
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

type User struct {
	*mongo.Collection
	model.Hashing
}

func (mongoDB *User) Save(user *user.User) error {
	new := UserDocument(*user.ToPrimitive())

	hashed, err := mongoDB.Hashing.Hash(new.Password)

	if err != nil {
		return errors.BubbleUp(err, "Save")
	}

	new.Password = hashed

	_, err = mongoDB.Collection.InsertOne(context.Background(), &new)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(mongodb.HandleDuplicateKeyError(err), "Save")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a user",
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
			What:  "failure to verify a user",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Update(user *user.User) error {
	updated := UserDocument(*user.ToPrimitive())

	filter := bson.D{{Key: "id", Value: user.Id.Value}}

	hashed, err := mongoDB.Hashing.Hash(user.Password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	updated.Password = hashed

	_, err = mongoDB.Collection.ReplaceOne(context.Background(), filter, &updated)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a user",
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
			What:  "failure to delete a user",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (mongoDB *User) Search(criteria *model.RepositorySearchCriteria) (*user.User, error) {
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
			What:  "failure to decode a result",
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
			What:  "failure to create an user from a primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Index":     index,
			},
			Who: err,
		})
	}

	return found, nil
}

func NewUser(mongoDB *mongodb.MongoDB, name string, hashing model.Hashing) (model.Repository, error) {
	collection := mongoDB.Database.Collection(name)

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
			Where: "NewUser",
			What:  "failure to create indexes for user collection",
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
