package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func BubbleWithoutWhere() *Bubble {
	return &Bubble{
		What: services.Create.LoremIpsumWord(),
	}
}

func BubbleWithoutWhat() *Bubble {
	return &Bubble{
		Where: services.Create.LoremIpsumWord(),
	}
}
