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

func (unit *operationUnitSuite) Multiply() {
	unit.Run("Return multiplication", func() {
		sut := math.NewOperations()

		result := sut.Multiply(1, 1)

		require.NotEmpty(unit.T(), result)
	})
}

func TestOperationUnitSuite(t *testing.T) {
	suite.Run(t, new(operationUnitSuite))
}
