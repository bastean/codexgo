package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/collection"
)

const (
	CollectionName = "users"
)

var (
	OpenCollection = collection.OpenUser
)
