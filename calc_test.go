package calc

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		a, b int // Name is equation
		want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 2, b: 3, want: 5},
		{a: -2, b: -3, want: -5},
		// TODO -- overflow/underflow not handled
		// {a: math.MaxInt, b: 1, want: ??? }
		// {a: math.MinInt, b: -1, want: ??? }
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%d + %d = %d", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			// Setup
			addition := &Addition{}

			// Run test
			result := addition.Calculate(tc.a, tc.b)

			// Check results
			if result != tc.want {
				t.Errorf("want: %v, got: %v", tc.want, result)
			}
		})
	}
}
