package operations

type Operation interface {
	Perform(a, b float64) float64
}