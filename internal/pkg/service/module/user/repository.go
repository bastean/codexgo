package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb"
)

const (
	CollectionName = "users"
)

var (
	OpenCollection = mongodb.OpenCollection
)
