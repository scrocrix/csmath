package collatz

type Conjecture struct {
	Sequence []int
}

// NewConjecture constructs a new instance of Conjecture.
func NewConjecture() *Conjecture {
	return &Conjecture{}
}

// CollatzConjecture executes the conjecture named after Lothat Collatz, that given a positive "n", we can generate
// a sequence resulting in "1".
func (c *Conjecture) Run(i int) {
	// we make sure that the first number of the sequence gets added to it.
	if len(c.Sequence) == 0 {
		c.Sequence = append(c.Sequence, i)
	}

	if i != 1 {
		if i%2 == 0 {
			c.Sequence = append(c.Sequence, i/2)
			c.Run(i / 2)
		} else {
			c.Sequence = append(c.Sequence, 3*i+1)
			c.Run(3*i + 1)
		}
	}
}
