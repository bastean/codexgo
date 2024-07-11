package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/collection"
)

func UserCollection(mongoDB *mongodb.MongoDB, name string, hashing model.Hashing) (model.Repository, error) {
	collection, err := collection.NewUser(mongoDB, name, hashing)

	if err != nil {
		return nil, errors.BubbleUp(err, "UserCollection")
	}

	return collection, nil
}
