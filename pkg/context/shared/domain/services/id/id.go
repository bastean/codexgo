package id

import (
	"github.com/google/uuid"
)

const RExID = `^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`

func New() string {
	return uuid.NewString()
}
