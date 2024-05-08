package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomId() (model.ValueObject[string], error) {
	return valueObject.NewId(mother.Create.UUID())
}

func InvalidId() (model.ValueObject[string], error) {
	return valueObject.NewId("x")
}

func EmptyId() (model.ValueObject[string], error) {
	return valueObject.NewId("")
}
