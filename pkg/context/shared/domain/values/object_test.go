package values_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type ObjectTestSuite struct {
	suite.Suite
	SUT *values.Object[string]
}

func (s *ObjectTestSuite) SetupTest() {
	s.SUT = new(values.Object[string])
}

func (s *ObjectTestSuite) TestSetCreated() {
	s.NotPanics(func() { s.SUT.SetCreated(services.Create.TimeNow()) })
}

func (s *ObjectTestSuite) TestSetCreatedErrAlreadyDefined() {
	s.NotPanics(func() { s.SUT.SetCreated(services.Create.TimeNow()) })

	expected := "(SetCreated): Created is already set"

	s.PanicsWithValue(expected, func() { s.SUT.SetCreated(services.Create.TimeNow()) })
}

func (s *ObjectTestSuite) TestSetUpdated() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	s.NotPanics(func() { s.SUT.SetUpdated(services.Create.TimeRandomAfter(date)) })
}

func (s *ObjectTestSuite) TestSetUpdatedErrCreatedUndefined() {
	expected := "(SetUpdated): Created is not defined"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdated(services.Create.TimeNow()) })
}

func (s *ObjectTestSuite) TestSetUpdatedErrBeforeCreated() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	expected := "(SetUpdated): Time updated cannot be before created"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdated(services.Create.TimeRandomBefore(date)) })
}

func (s *ObjectTestSuite) TestSetUpdatedErrEqualsCreated() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	expected := "(SetUpdated): Time updated cannot be equal to created"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdated(date) })
}

func (s *ObjectTestSuite) TestSetUpdatedErrBeforeDefined() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	s.NotPanics(func() { s.SUT.SetUpdated(services.Create.TimeSetAfter(date, 48, 72)) })

	expected := "(SetUpdated): Updated time cannot be before existing value"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdated(services.Create.TimeSetAfter(date, 1, 24)) })
}

func (s *ObjectTestSuite) TestSet() {
	s.NotPanics(func() { s.SUT.Set(services.Create.LoremIpsumWord()) })
}

func (s *ObjectTestSuite) TestSetErrAlreadyDefined() {
	s.NotPanics(func() { s.SUT.Set(services.Create.LoremIpsumWord()) })

	expected := "(Set): Value is already set"

	s.PanicsWithValue(expected, func() { s.SUT.Set(services.Create.LoremIpsumWord()) })
}

func (s *ObjectTestSuite) TestValid() {
	s.NotPanics(func() { s.SUT.Set(services.Create.LoremIpsumWord()) })

	s.SUT.Valid()

	s.NotPanics(func() { _ = s.SUT.Value() })
}

func (s *ObjectTestSuite) TestValidErrAlreadyValidated() {
	s.NotPanics(func() { s.SUT.Set(services.Create.LoremIpsumWord()) })

	s.SUT.Valid()

	expected := "(Valid): Value is already validated"

	s.PanicsWithValue(expected, func() { s.SUT.Valid() })
}
func (s *ObjectTestSuite) TestValidErrNoValue() {
	expected := "(Valid): No value to validate"
	s.PanicsWithValue(expected, func() { s.SUT.Valid() })
}

func (s *ObjectTestSuite) TestValidateErrValidationNotDefined() {
	err := s.SUT.Validate()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Validate",
		What:  "Value validation is not defined",
	}}

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValidateWithCustomValidation() {
	object := values.Mother().ObjectWithCustomValidation()

	s.NotPanics(func() { object.Set("") })

	err := object.Validate()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Validate",
		What:  "Value can not be empty",
	}}

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValue() {
	expected := services.Create.LoremIpsumWord()

	s.NotPanics(func() { s.SUT.Set(expected) })

	s.SUT.Valid()

	actual := s.SUT.Value()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValueErrNotValidated() {
	expected := "(Value): Value needs to be validated"
	s.PanicsWithValue(expected, func() { _ = s.SUT.Value() })
}

func (s *ObjectTestSuite) TestRawValue() {
	expected := services.Create.LoremIpsumWord()

	s.NotPanics(func() { s.SUT.Set(expected) })

	actual := s.SUT.RawValue()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestCreated() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	actual := s.SUT.Created()

	expected := date.UTC().Format(time.RFC3339Nano)

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestUpdated() {
	date := services.Create.TimeNow()

	s.NotPanics(func() { s.SUT.SetCreated(date) })

	date = services.Create.TimeRandomAfter(date)

	s.NotPanics(func() { s.SUT.SetUpdated(date) })

	actual := s.SUT.Updated()

	expected := date.UTC().Format(time.RFC3339Nano)

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestNew() {
	actual, err := values.New[*values.Example](services.Create.LoremIpsumWord())

	s.NoError(err)

	var expected *values.Example

	s.IsType(expected, actual)
}

func (s *ObjectTestSuite) TestReplace() {
	actual, err := values.New[*values.Example](services.Create.LoremIpsumWord())

	s.NoError(err)

	expected, err := values.Replace(actual, services.Create.LoremIpsumWord())

	s.NoError(err)

	s.NotEqual(expected.Value(), actual.Value())

	s.Equal(expected.Created(), actual.Created())

	s.NotEqual(expected.Updated(), actual.Updated())
}

func TestUnitObjectSuite(t *testing.T) {
	suite.Run(t, new(ObjectTestSuite))
}
