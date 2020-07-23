package operations

type Mul struct{}

func (a2 Mul) Perform(a, b float64) float64 {
	return a * b
}
