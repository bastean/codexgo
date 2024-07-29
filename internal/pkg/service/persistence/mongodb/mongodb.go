package mongodb

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
)

type MongoDB = mongodb.MongoDB

var (
	Open  = mongodb.Open
	Close = mongodb.Close
)
