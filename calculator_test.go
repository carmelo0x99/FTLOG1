package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{a: 2, b: 2, want: 3, name: "2 + 2 = 4"},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			if tc.name != "" {
				fmt.Println("[!] Failed test: ", tc.name)
			}
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []testCase{
		{a: 101, b: 100, want: 1},
		{a: 2, b: 4, want: -2},
		{a: 10, b: 3, want: 7},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []testCase{
		{a: 6, b: 7, want: 42},
		{a: 1000, b: 0, want: 0},
		{a: 3, b: -3, want: -9},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
