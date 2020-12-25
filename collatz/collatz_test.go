package collatz_test

import (
	"github.com/scrocrix/csmath/collatz"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type conjectureUnitTest struct {
	suite.Suite
}

func (unit *conjectureUnitTest) TestRun() {
	unit.Run("Return a slice of integers as sequenece", func() {
		sut := collatz.NewConjecture()

		sut.Run(3)

		require.Equal(unit.T(), []int{3, 10, 5, 16, 8, 4, 2, 1}, sut.Sequence)
	})
}

func TestConjectureUnitTest(t *testing.T) {
	suite.Run(t, new(conjectureUnitTest))
}
