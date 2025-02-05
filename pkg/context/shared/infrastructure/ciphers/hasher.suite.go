package ciphers

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type HasherSuite struct {
	suite.Suite
	SUT roles.Hasher
}

func (s *HasherSuite) TestHash() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *HasherSuite) TestCompare() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	err = s.SUT.Compare(hashed, plain)

	s.NoError(err)
}

func (s *HasherSuite) TestCompareErrDoNotMatch() {
	plain := services.Create.LoremIpsumWord()

	hashed := services.Create.LoremIpsumWord()

	err := s.SUT.Compare(hashed, plain)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Compare",
		What:  "Password does not match",
	}}

	s.Equal(expected, actual)
}
