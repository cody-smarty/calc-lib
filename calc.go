package calc

type (
	Addition       struct{}
	Subtraction    struct{}
	Multiplication struct{}
	Division       struct{}
)

func (add *Addition) Calculate(a, b int) int       { return a + b }
func (add *Subtraction) Calculate(a, b int) int    { return a - b }
func (add *Multiplication) Calculate(a, b int) int { return a * b }
func (add *Division) Calculate(a, b int) int       { return a / b } // No divide-by-zero protection
