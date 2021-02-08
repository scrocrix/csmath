package fibonacci_test

import (
	"github.com/scrocrix/csmath/fibonacci"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type fibonacciUnitTest struct {
	suite.Suite
}

func (unit *fibonacciUnitTest) TestNew() {
	unit.Run("Return a new instance of fibonacci", func() {
		sut := fibonacci.New()

		require.NotNil(unit.T(), sut)
	})
}

func (unit *fibonacciUnitTest) TestFindTerm() {
	unit.Run("Return nth term of the sequence", func() {
		sut := fibonacci.New()
		require.Equal(unit.T(), 5, sut.FindTerm(4))
		require.Equal(unit.T(), 8, sut.FindTerm(5))
	})
}

func TestFibonacciUnitTest(t *testing.T) {
	suite.Run(t, new(fibonacciUnitTest))
}
