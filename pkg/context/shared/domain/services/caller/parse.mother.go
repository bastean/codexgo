package caller

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func ParseWithRandomValues() (string, string, string) {
	return services.Create.LoremIpsumWord(), services.Create.LoremIpsumWord(), services.Create.LoremIpsumWord()
}
