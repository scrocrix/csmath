package binets_test

import (
	"github.com/scrocrix/csmath/binets"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type binetsUnitTest struct {
	suite.Suite
}

func (unit *binetsUnitTest) TestNewBinets() {
	unit.Run("Return a new instance of Binets", func() {
		sut := binets.New()

		require.NotNil(unit.T(), sut)
	})
}

func (unit *binetsUnitTest) TestCalculate() {
	unit.Run("Return fibonacci number matching the position", func() {
		sut := binets.New()

		require.Equal(unit.T(), 1, sut.Calculate(1))
		require.Equal(unit.T(), 1, sut.Calculate(2))
		require.Equal(unit.T(), 2, sut.Calculate(3))
		require.Equal(unit.T(), 3, sut.Calculate(4))
		require.Equal(unit.T(), 5, sut.Calculate(5))
		require.Equal(unit.T(), 8, sut.Calculate(6))
		require.Equal(unit.T(), 13, sut.Calculate(7))
	})
}

func TestBinetsUnitTest(t *testing.T) {
	suite.Run(t, new(binetsUnitTest))
}
