package calc

type Addition struct{}

func (add *Addition) Calculate(a, b int) int {
	return a + b
}
