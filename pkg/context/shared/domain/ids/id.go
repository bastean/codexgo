package ids

import (
	"github.com/google/uuid"
)

func Generate() string {
	return uuid.NewString()
}
