package binets

import (
	"math"
	"math/big"
)

type Binets interface {
	// Calculate return the fibonacci number in the n-th position argument. (e.g. 6th fibonacci number is 8)
	Calculate(position int) int
}

type binets struct {
	Binets
}

// New constructs a new instance of Binets.
func New() Binets {
	return &binets{}
}

func (b *binets) calculateGoldenRatio(isInverse bool) float64 {
	if isInverse {
		return (1 - math.Sqrt(float64(5))) / 2
	}

	return (1 + math.Sqrt(float64(5))) / 2
}

func (b *binets) Calculate(position int) int {
	f := big.NewFloat(math.Pow(b.calculateGoldenRatio(false), float64(position)))
	g := big.NewFloat(math.Pow(b.calculateGoldenRatio(true), float64(position)))

	result := new(big.Float).Sub(f, g)

	ff, _ := result.Float64()

	return int(float64(1/math.Sqrt(5)) * ff)
}
