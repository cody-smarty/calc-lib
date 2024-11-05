package calc

import (
	"fmt"
	"testing"
)

func TestCalculate_Addition(t *testing.T) {
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

func TestCalculate_Subtraction(t *testing.T) {
	testCases := []struct {
		a, b int // Name is equation
		want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 2, b: 3, want: -1},
		{a: -2, b: -3, want: 1},
		// TODO -- overflow/underflow not handled
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%d + %d = %d", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			// Setup
			subtraction := &Subtraction{}

			// Run test
			result := subtraction.Calculate(tc.a, tc.b)

			// Check results
			if result != tc.want {
				t.Errorf("want: %v, got: %v", tc.want, result)
			}
		})
	}
}

func TestCalculate_Multiplication(t *testing.T) {
	testCases := []struct {
		a, b int // Name is equation
		want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 2, b: 3, want: 6},
		{a: -2, b: -3, want: 6},
		// TODO -- overflow/underflow not handled
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%d + %d = %d", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			// Setup
			multiplication := &Multiplication{}

			// Run test
			result := multiplication.Calculate(tc.a, tc.b)

			// Check results
			if result != tc.want {
				t.Errorf("want: %v, got: %v", tc.want, result)
			}
		})
	}
}

func TestCalculate_Division(t *testing.T) {
	testCases := []struct {
		a, b int // Name is equation
		want int
	}{
		{a: 0, b: 1, want: 0},
		{a: 2, b: 3, want: 0},
		{a: -2, b: -3, want: 0},
		// TODO -- overflow/underflow not handled
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%d + %d = %d", tc.a, tc.b, tc.want)
		t.Run(testName, func(t *testing.T) {
			// Setup
			division := &Division{}

			// Run test
			result := division.Calculate(tc.a, tc.b)

			// Check results
			if result != tc.want {
				t.Errorf("want: %v, got: %v", tc.want, result)
			}
		})
	}
}

func TestDivideByZeroPanics(t *testing.T) {
	division := &Division{}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	division.Calculate(1, 0)
}
