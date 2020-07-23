package operations

type Add struct{}

func (a2 Add) Perform(a, b float64) float64 {
	return a + b
}
