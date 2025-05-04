package id_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/id"
)

type IDTestSuite struct {
	suite.Suite
}

func (s *IDTestSuite) SetupSuite() {
	s.Equal(id.RExID, `^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`)
}

func (s *IDTestSuite) TestNew() {
	s.Regexp(id.RExID, id.New())
}

func (s *IDTestSuite) TestNewUnique() {
	s.NotEqual(id.New(), id.New())
}

func TestUnitIDSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}
