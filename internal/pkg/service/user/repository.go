package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/collection"
)

func Collection(mongoDB *mongodb.MongoDB, name string, hashing hashing.Hashing) (repository.User, error) {
	collection, err := collection.NewUser(mongoDB, name, hashing)

	if err != nil {
		return nil, errors.BubbleUp(err, "Collection")
	}

	return collection, nil
}
