package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
)

const (
	CollectionName = "users"
)

var (
	OpenCollection = collection.Open
)
