package loggers

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomMessage() string {
	return services.Create.LoremIpsumSentence(services.Create.RandomInt([]int{1, 10}))
}
