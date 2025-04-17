package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func BubbleUpWithRandomValue() (error, error) {
	err := services.Create.Error()
	return BubbleUp(err), err
}
