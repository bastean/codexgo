package values_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Custom struct {
	values.String
}

func (c *Custom) Validate() error {
	if c.RawValue() == "" {
		return errors.New[errors.InvalidValue](&errors.Bubble{
			What: "Value can not be empty",
		})
	}

	c.Valid()

	return nil
}

type ObjectTestSuite struct {
	suite.Default
	SUT *values.String
}

func (s *ObjectTestSuite) SetupTest() {
	s.SUT = new(values.String)
}

func (s *ObjectTestSuite) TestSetCreatedAt() {
	s.NotPanics(func() { s.SUT.SetCreatedAt(values.Mother().TimeNow()) })
}

func (s *ObjectTestSuite) TestSetCreatedAtErrAlreadySet() {
	s.NotPanics(func() { s.SUT.SetCreatedAt(values.Mother().TimeNow()) })

	expected := "(values/*Object/SetCreatedAt): Created is already set"

	s.PanicsWithValue(expected, func() { s.SUT.SetCreatedAt(values.Mother().TimeNow()) })
}

func (s *ObjectTestSuite) TestSetUpdatedAt() {
	date := values.Mother().TimeNow()

	s.NotPanics(func() { s.SUT.SetCreatedAt(date) })

	s.NotPanics(func() { s.SUT.SetUpdatedAt(values.Mother().TimeRandomAfter(date)) })
}

func (s *ObjectTestSuite) TestSetUpdatedAtErrCreatedNotDefined() {
	expected := "(values/*Object/SetUpdatedAt): Created is not defined"
	s.PanicsWithValue(expected, func() { s.SUT.SetUpdatedAt(values.Mother().TimeNow()) })
}

func (s *ObjectTestSuite) TestSetUpdatedAtErrCannotBeBeforeCreated() {
	date := values.Mother().TimeNow()

	s.NotPanics(func() { s.SUT.SetCreatedAt(date) })

	expected := "(values/*Object/SetUpdatedAt): Time updated cannot be before created"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdatedAt(values.Mother().TimeRandomBefore(date)) })
}

func (s *ObjectTestSuite) TestSetUpdatedAtErrCannotBeBeforeExistingValue() {
	date := values.Mother().TimeNow()

	s.NotPanics(func() { s.SUT.SetCreatedAt(date) })

	s.NotPanics(func() { s.SUT.SetUpdatedAt(values.Mother().TimeSetAfter(date, time.Day*2, time.Day*3)) })

	expected := "(values/*Object/SetUpdatedAt): Updated time cannot be before existing value"

	s.PanicsWithValue(expected, func() { s.SUT.SetUpdatedAt(values.Mother().TimeSetAfter(date, time.Hour, time.Day)) })
}

func (s *ObjectTestSuite) TestSet() {
	s.NotPanics(func() { s.SUT.Set(values.Mother().LoremIpsumWord()) })
}

func (s *ObjectTestSuite) TestSetErrAlreadySet() {
	s.NotPanics(func() { s.SUT.Set(values.Mother().LoremIpsumWord()) })

	expected := "(values/*Object/Set): Value is already set"

	s.PanicsWithValue(expected, func() { s.SUT.Set(values.Mother().LoremIpsumWord()) })
}

func (s *ObjectTestSuite) TestValid() {
	s.NotPanics(func() { s.SUT.Set(values.Mother().LoremIpsumWord()) })

	s.SUT.Valid()

	s.NotPanics(func() { _ = s.SUT.Value() })
}

func (s *ObjectTestSuite) TestValidErrAlreadyValidated() {
	s.NotPanics(func() { s.SUT.Set(values.Mother().LoremIpsumWord()) })

	s.SUT.Valid()

	expected := "(values/*Object/Valid): Value is already validated"

	s.PanicsWithValue(expected, func() { s.SUT.Valid() })
}
func (s *ObjectTestSuite) TestValidErrNoValue() {
	expected := "(values/*Object/Valid): No value to validate"
	s.PanicsWithValue(expected, func() { s.SUT.Valid() })
}

func (s *ObjectTestSuite) TestValidateErrValidationNotDefined() {
	err := s.SUT.Validate()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "values/*Object/Validate",
		What:  "Value validation is not defined",
	}}

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValidateWithCustomValidation() {
	object := new(Custom)

	s.NotPanics(func() { object.Set("") })

	err := object.Validate()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "values_test/*Custom/Validate",
		What:  "Value can not be empty",
	}}

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValue() {
	expected := values.Mother().LoremIpsumWord()

	s.NotPanics(func() { s.SUT.Set(expected) })

	s.SUT.Valid()

	actual := s.SUT.Value()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestValueErrNeedsToBeValidated() {
	expected := "(values/*Object/Value): Value needs to be validated"
	s.PanicsWithValue(expected, func() { _ = s.SUT.Value() })
}

func (s *ObjectTestSuite) TestRawValue() {
	expected := values.Mother().LoremIpsumWord()

	s.NotPanics(func() { s.SUT.Set(expected) })

	actual := s.SUT.RawValue()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestCreatedAt() {
	date := values.Mother().TimeNow()

	s.NotPanics(func() { s.SUT.SetCreatedAt(date) })

	actual := s.SUT.CreatedAt()

	expected := date.Format()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestUpdatedAt() {
	date := values.Mother().TimeNow()

	s.NotPanics(func() { s.SUT.SetCreatedAt(date) })

	date = values.Mother().TimeRandomAfter(date)

	s.NotPanics(func() { s.SUT.SetUpdatedAt(date) })

	actual := s.SUT.UpdatedAt()

	expected := date.Format()

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestNew() {
	actual, err := values.New[*Custom](values.Mother().LoremIpsumWord())

	s.NoError(err)

	var expected *Custom

	s.IsType(expected, actual)
}
func (s *ObjectTestSuite) TestFromPrimitive() {
	expected, err := values.New[*Custom](values.Mother().LoremIpsumWord())

	s.NoError(err)

	actual, err := values.FromPrimitive[*Custom](expected.ToPrimitive())

	s.NoError(err)

	s.IsType(expected, actual)

	s.Equal(expected.Value(), actual.Value())

	s.Equal(expected.CreatedAt(), actual.CreatedAt())

	s.Equal(expected.UpdatedAt(), actual.UpdatedAt())
}

func (s *ObjectTestSuite) TestFromPrimitiveWithOptional() {
	actual, err := values.FromPrimitive[*Custom](nil, true)

	s.NoError(err)

	var expected *Custom

	s.IsType(expected, actual)
}

func (s *ObjectTestSuite) TestFromPrimitiveErrValueRequired() {
	_, err := values.FromPrimitive[*Custom](nil)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "values/FromPrimitive",
		What:  "Primitive value is required",
	}}

	s.Equal(expected, actual)
}

func (s *ObjectTestSuite) TestReplace() {
	actual, err := values.New[*Custom](values.Mother().LoremIpsumWord())

	s.NoError(err)

	expected, err := values.Replace(actual, values.Mother().Message())

	s.NoError(err)

	s.NotEqual(expected.Value(), actual.Value())

	s.Equal(expected.CreatedAt(), actual.CreatedAt())

	s.NotEqual(expected.UpdatedAt(), actual.UpdatedAt())
}

func TestUnitObjectSuite(t *testing.T) {
	suite.Run(t, new(ObjectTestSuite))
}
