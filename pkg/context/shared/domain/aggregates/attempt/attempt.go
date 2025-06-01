package attempt

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Attempt struct {
	Counter, Limit *values.IntPositive
	Until          *values.Time
	Every, Next    *values.IntPositive
}

type Primitive struct {
	Counter, Limit *values.IntPrimitive
	Until          *values.StringPrimitive
	Every, Next    *values.IntPrimitive
}

type Required struct {
	Limit       int
	Every, Next int
}

func (a *Attempt) Increase() error {
	current := time.Now()

	switch {
	case a.Until != nil && current.Before(time.Parse(a.Until.Value())):
		return errors.New[errors.Failure](&errors.Bubble{
			What: fmt.Sprintf("No more attempts, please try again in %q", current.Sub(time.Parse(a.Until.Value())).Round(time.Second)),
		})
	case a.Counter.Updated() != "" && current.Before(time.Parse(a.Counter.Updated()).Add(time.Duration(a.Every.Value()))):
		return errors.New[errors.Failure](&errors.Bubble{
			What: fmt.Sprintf("Please try again in %q", current.Sub(time.Parse(a.Counter.Updated()).Add(time.Duration(a.Every.Value()))).Round(time.Second)),
		})
	}

	counter, err := values.Replace(a.Counter, a.Counter.Value()+1)

	if err != nil {
		return errors.BubbleUp(err)
	}

	if counter.Value() == a.Limit.Value() {
		counter, err = values.Replace(a.Counter, 0)

		if err != nil {
			return errors.BubbleUp(err)
		}

		var until *values.Time

		switch a.Until {
		case nil:
			until, err = values.New[*values.Time](current.Add(time.Duration(a.Next.Value())).Format())
		default:
			until, err = values.Replace(a.Until, current.Add(time.Duration(a.Next.Value())).Format())
		}

		if err != nil {
			return errors.BubbleUp(err)
		}

		a.Until = until
	}

	a.Counter = counter

	return nil
}

func (a *Attempt) ToPrimitive() *Primitive {
	primitive := &Primitive{
		Counter: a.Counter.ToPrimitive(),
		Limit:   a.Limit.ToPrimitive(),
		Every:   a.Every.ToPrimitive(),
		Next:    a.Next.ToPrimitive(),
	}

	if a.Until != nil {
		primitive.Until = a.Until.ToPrimitive()
	}

	return primitive
}

func FromPrimitive(primitive *Primitive) (*Attempt, error) {
	counter, errCounter := values.FromPrimitive[*values.IntPositive](primitive.Counter)
	limit, errLimit := values.FromPrimitive[*values.IntPositive](primitive.Limit)

	until, errUntil := values.FromPrimitive[*values.Time](primitive.Until, true)

	every, errEvery := values.FromPrimitive[*values.IntPositive](primitive.Every)
	next, errNext := values.FromPrimitive[*values.IntPositive](primitive.Next)

	if err := errors.Join(errCounter, errLimit, errUntil, errEvery, errNext); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return &Attempt{
		Counter: counter,
		Limit:   limit,
		Until:   until,
		Every:   every,
		Next:    next,
	}, nil
}

func New(required *Required) (*Attempt, error) {
	counter, errCounter := values.New[*values.IntPositive](0)
	limit, errLimit := values.New[*values.IntPositive](required.Limit)

	every, errEvery := values.New[*values.IntPositive](required.Every)
	next, errNext := values.New[*values.IntPositive](required.Next)

	if err := errors.Join(errCounter, errLimit, errEvery, errNext); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return &Attempt{
		Counter: counter,
		Limit:   limit,
		Every:   every,
		Next:    next,
	}, nil
}
