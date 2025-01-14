package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Collection struct {
	*mongo.Collection
}

func (c *Collection) Create(user *user.User) error {
	err := user.CreationStamp()

	if err != nil {
		return errors.BubbleUp(err, "Create")
	}

	aggregate := user.ToPrimitive()

	_, err = c.Collection.InsertOne(context.Background(), aggregate)

	switch {
	case mongodb.IsErrDuplicateValue(err):
		return errors.BubbleUp(mongodb.HandleErrDuplicateValue(err), "Create")
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Create",
			What:  "Failure to create a User",
			Why: errors.Meta{
				"ID": user.ID.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (c *Collection) Update(user *user.User) error {
	err := user.UpdatedStamp()

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	aggregate := user.ToPrimitive()

	filter := bson.D{{Key: "id", Value: user.ID.Value}}

	_, err = c.Collection.ReplaceOne(context.Background(), filter, aggregate)

	switch {
	case mongodb.IsErrDuplicateValue(err):
		return errors.BubbleUp(mongodb.HandleErrDuplicateValue(err), "Update")
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Update",
			What:  "Failure to update a User",
			Why: errors.Meta{
				"ID": user.ID.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (c *Collection) Delete(id *user.ID) error {
	filter := bson.D{{Key: "id", Value: id.Value}}

	_, err := c.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Delete",
			What:  "Failure to delete a User",
			Why: errors.Meta{
				"ID": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (c *Collection) Search(criteria *user.Criteria) (*user.User, error) {
	var (
		filter bson.D
		index  string
	)

	switch {
	case criteria.ID != nil:
		filter = bson.D{{Key: "id", Value: criteria.ID.Value}}
		index = criteria.ID.Value
	case criteria.Email != nil:
		filter = bson.D{{Key: "email", Value: criteria.Email.Value}}
		index = criteria.Email.Value
	case criteria.Username != nil:
		filter = bson.D{{Key: "username", Value: criteria.Username.Value}}
		index = criteria.Username.Value
	default:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Criteria not defined",
		})
	}

	result := c.Collection.FindOne(context.Background(), filter)

	err := result.Err()

	switch {
	case mongodb.IsErrNotFound(err):
		return nil, mongodb.HandleErrNotFound(err, index)
	case err != nil:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Failure to find a User",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	primitive := new(user.Primitive)

	err = result.Decode(primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Failure to decode a result",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	aggregate, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Search",
			What:  "Failure to create a User from a Primitive",
			Why: errors.Meta{
				"Index":     index,
				"Primitive": primitive,
			},
			Who: err,
		})
	}

	return aggregate, nil
}

func Open(session *mongodb.Database, name string) (role.Repository, error) {
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
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure to create Indexes for User Collection",
			Why: errors.Meta{
				"Collection": name,
			},
			Who: err,
		})
	}

	return &Collection{
		Collection: collection,
	}, nil
}
