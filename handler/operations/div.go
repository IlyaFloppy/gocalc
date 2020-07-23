package operations

type Div struct{}

func (a2 Div) Perform(a, b float64) float64 {
	return a / b
}
