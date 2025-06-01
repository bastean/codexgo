package values

import (
	"reflect"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type Primitive[T any] struct {
	Value            T
	Created, Updated string
}

type valueObject[T any] interface {
	SetCreated(time.Time)
	SetUpdated(time.Time)
	Set(T)
	Valid()
	Validate() error
	Value() T
	RawValue() T
	Created() string
	Updated() string
	ToPrimitive() *Primitive[T]
}

type Object[T any] struct {
	value            T
	isSet, isValid   bool
	created, updated time.Time
}

func (o *Object[T]) SetCreated(actual time.Time) {
	if !o.created.IsZero() {
		errors.Panic(errors.Standard("Created is already set"))
	}

	o.created = actual
}

func (o *Object[T]) SetUpdated(actual time.Time) {
	switch {
	case o.created.IsZero():
		errors.Panic(errors.Standard("Created is not defined"))
	case actual.Before(o.created):
		errors.Panic(errors.Standard("Time updated cannot be before created"))
	case actual.Before(o.updated):
		errors.Panic(errors.Standard("Updated time cannot be before existing value"))
	}

	o.updated = actual
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

func (o *Object[T]) Created() string {
	return o.created.Format()
}

func (o *Object[T]) Updated() string {
	if !o.updated.IsZero() {
		return o.updated.Format()
	}

	return ""
}

func (o *Object[T]) ToPrimitive() *Primitive[T] {
	return &Primitive[T]{
		Value:   o.Value(),
		Created: o.Created(),
		Updated: o.Updated(),
	}
}

func create[O valueObject[V], V any](value V) (O, error) {
	object := reflect.New(reflect.TypeOf(*new(O)).Elem()).Interface().(O)

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

	object.SetCreated(time.Now())

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

	object.SetCreated(time.Parse(primitive.Created))

	if primitive.Updated != "" {
		object.SetUpdated(time.Parse(primitive.Updated))
	}

	return object, nil
}

func Replace[O valueObject[V], V any](old O, value V) (O, error) {
	object, err := create[O](value)

	if err != nil {
		return *new(O), errors.BubbleUp(err)
	}

	object.SetCreated(time.Parse(old.Created()))

	object.SetUpdated(time.Now())

	return object, nil
}
