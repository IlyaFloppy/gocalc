package operations

type Sub struct{}

func (a2 Sub) Perform(a, b float64) float64 {
	return a - b
}
