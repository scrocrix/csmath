package math

type Operations interface {
	// Multiply returns the product of two terms
	Multiply(termOne int, termTwo int) int
}

type operations struct {
	Operations
}

// NewOperations constructs a new instance of Operations.
func NewOperations() Operations {
	return &operations{}
}

func (o *operations) Multiply(termOne int, termTwo int) int {
	return termOne * termTwo
}
