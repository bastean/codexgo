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
	_, err := c.Collection.InsertOne(context.Background(), user.ToPrimitive())

	switch {
	case mongodb.IsErrDuplicateValue(err):
		return errors.BubbleUp(mongodb.HandleErrDuplicateValue(err))
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to create a User",
			Why: errors.Meta{
				"ID": user.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (c *Collection) Update(user *user.User) error {
	_, err := c.Collection.ReplaceOne(context.Background(),
		bson.D{{Key: "id.value", Value: user.ID.Value()}},
		user.ToPrimitive(),
	)

	switch {
	case mongodb.IsErrDuplicateValue(err):
		return errors.BubbleUp(mongodb.HandleErrDuplicateValue(err))
	case err != nil:
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to update a User",
			Why: errors.Meta{
				"ID": user.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (c *Collection) Delete(id *user.ID) error {
	_, err := c.Collection.DeleteOne(context.Background(),
		bson.D{{Key: "id.value", Value: id.Value()}},
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to delete a User",
			Why: errors.Meta{
				"ID": id.Value(),
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
		filter = bson.D{{Key: "id.value", Value: criteria.ID.Value()}}
		index = criteria.ID.Value()
	case criteria.Email != nil:
		filter = bson.D{{Key: "email.value", Value: criteria.Email.Value()}}
		index = criteria.Email.Value()
	case criteria.Username != nil:
		filter = bson.D{{Key: "username.value", Value: criteria.Username.Value()}}
		index = criteria.Username.Value()
	default:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Criteria not defined",
		})
	}

	result := c.Collection.FindOne(context.Background(), filter)

	err := result.Err()

	switch {
	case mongodb.IsErrNotFound(err):
		return nil, mongodb.HandleErrNotFound(err, index)
	case err != nil:
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to find a User",
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
			What: "Failure to decode a result",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	aggregate, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to create a User from a Primitive",
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
			Keys:    bson.D{{Key: "id.value", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email.value", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username.value", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to create Indexes for User Collection",
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
