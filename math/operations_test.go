package math_test

import (
	"github.com/scrocrix/csmath/math"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type operationUnitSuite struct {
	suite.Suite
}

func (unit *operationUnitSuite) TestComplexMultiply() {
	unit.Run("Return multiplication", func() {
		sut := math.NewOperations()

		result := sut.ComplexMultiply(1, 1, 2)

		require.NotEmpty(unit.T(), result)

		require.Equal(unit.T(), 1, result)
	})

	unit.Run("Return zero when second term is equal to zero", func() {
		sut := math.NewOperations()

		require.Equal(unit.T(), 0, sut.ComplexMultiply(1, 0, 2))
	})
}

func TestOperationUnitSuite(t *testing.T) {
	suite.Run(t, new(operationUnitSuite))
}
