package fibonacci

type Fibonacci interface {
	// FindTerm calculates the nth term in the Fibonacci sequence.
	//
	// The number gets calculated through Fibonacci's recursive formula: fn = fn-1 + fn-2
	FindTerm(n int) int
}

type fibonacci struct {
	Fibonacci
}

// New constructs a new instance of Fibonacci
func New() Fibonacci {
	return &fibonacci{}
}

func (f *fibonacci) FindTerm(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return f.FindTerm(n-1) + f.FindTerm(n-2)
}
