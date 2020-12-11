package math

type Operations interface {
	// ComplexMultiply returns the product of two terms.
	//
	// This method is considered "complex" because it takes a recursive approach rather than a straight forward
	// multiplication statement.
	ComplexMultiply(y int, z int, c int) int
}

type operations struct {
	Operations
}

// NewOperations constructs a new instance of Operations.
func NewOperations() Operations {
	return &operations{}
}

func (o *operations) ComplexMultiply(y int, z int, c int) int {
	if z == 0 {
		return 0
	}

	if c < 2 {
		c = 2
	}

	return o.ComplexMultiply(c*y, z/c, c) + y*(z%c)
}
