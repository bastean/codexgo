package values

import (
	"reflect"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type Primitive[T any] struct {
	Value                T
	CreatedAt, UpdatedAt string
}

type valueObject[T any] interface {
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	Set(T)
	Valid()
	Validate() error
	Value() T
	RawValue() T
	CreatedAt() string
	UpdatedAt() string
	ToPrimitive() *Primitive[T]
}

type Object[T any] struct {
	value                T
	isSet, isValid       bool
	createdAt, updatedAt time.Time
}

func (o *Object[T]) SetCreatedAt(actual time.Time) {
	if !o.createdAt.IsZero() {
		errors.Panic(errors.Standard("Created is already set"))
	}

	o.createdAt = actual
}

func (o *Object[T]) SetUpdatedAt(actual time.Time) {
	switch {
	case o.createdAt.IsZero():
		errors.Panic(errors.Standard("Created is not defined"))
	case actual.Before(o.createdAt):
		errors.Panic(errors.Standard("Time updated cannot be before created"))
	case actual.Before(o.updatedAt):
		errors.Panic(errors.Standard("Updated time cannot be before existing value"))
	}

	o.updatedAt = actual
}

func (o *Object[T]) Set(value T) {
	if o.isSet {
		errors.Panic(errors.Standard("Value is already set"))
	}

	o.value = value

	o.isSet = true
}

func (o *Object[T]) Valid() {
	switch {
	case o.isValid:
		errors.Panic(errors.Standard("Value is already validated"))
	case !o.isSet:
		errors.Panic(errors.Standard("No value to validate"))
	}

	o.isValid = true
}

func (*Object[T]) Validate() error {
	return errors.New[errors.Internal](&errors.Bubble{
		What: "Value validation is not defined",
	})
}

func (o *Object[T]) Value() T {
	if !o.isValid {
		errors.Panic(errors.Standard("Value needs to be validated"))
	}

	return o.value
}

func (o *Object[T]) RawValue() T {
	return o.value
}

func (o *Object[T]) CreatedAt() string {
	return o.createdAt.Format()
}

func (o *Object[T]) UpdatedAt() string {
	if !o.updatedAt.IsZero() {
		return o.updatedAt.Format()
	}

	return ""
}

func (o *Object[T]) ToPrimitive() *Primitive[T] {
	return &Primitive[T]{
		Value:     o.Value(),
		CreatedAt: o.CreatedAt(),
		UpdatedAt: o.UpdatedAt(),
	}
}

func create[O valueObject[V], V any](value V) (O, error) {
	object, _ := reflect.TypeAssert[O](reflect.New(reflect.TypeOf(*new(O)).Elem()))

	object.Set(value)

	if err := object.Validate(); err != nil {
		return *new(O), errors.BubbleUp(err)
	}

	return object, nil
}

func New[O valueObject[V], V any](value V) (O, error) {
	object, err := create[O](value)

	if err != nil {
		return *new(O), errors.BubbleUp(err)
	}

	object.SetCreatedAt(time.Now())

	return object, nil
}

func FromPrimitive[O valueObject[V], V any](primitive *Primitive[V], isOptional ...bool) (O, error) {
	if primitive == nil {
		switch {
		case len(isOptional) == 1:
			return *new(O), nil
		default:
			return *new(O), errors.New[errors.Internal](&errors.Bubble{
				What: "Primitive value is required",
			})
		}
	}

	object, err := create[O](primitive.Value)

	if err != nil {
		return *new(O), errors.BubbleUp(err)
	}

	object.SetCreatedAt(time.Parse(primitive.CreatedAt))

	if primitive.UpdatedAt != "" {
		object.SetUpdatedAt(time.Parse(primitive.UpdatedAt))
	}

	return object, nil
}

func Replace[O valueObject[V], V any](old O, value V) (O, error) {
	object, err := create[O](value)

	if err != nil {
		return *new(O), errors.BubbleUp(err)
	}

	object.SetCreatedAt(time.Parse(old.CreatedAt()))

	object.SetUpdatedAt(time.Now())

	return object, nil
}
