package ciphers

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type HasherSuite struct {
	suite.Default
	SUT roles.Hasher
}

func (s *HasherSuite) TestHash() {
	plain := Mother().LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *HasherSuite) TestCompare() {
	plain := Mother().LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	err = s.SUT.Compare(hashed, plain)

	s.NoError(err)
}

func (s *HasherSuite) TestCompareErrDoNotMatch() {
	plain := Mother().LoremIpsumWord()

	hashed := Mother().LoremIpsumWord()

	err := s.SUT.Compare(hashed, plain)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "Compare")

	expected := &errors.Failure{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: actual.Where,
		What:  "Password does not match",
	}}

	s.Equal(expected, actual)
}
