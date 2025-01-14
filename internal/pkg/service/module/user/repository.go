package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/sqlite/table"
)

type (
	Repository = role.Repository
)

const (
	CollectionName = "users"
)

var (
	OpenCollection = collection.Open
	OpenTable      = table.Open
)
