Let's now improve upin our test strategy by making the tests more modular. With a `struct`, such as:<br/>
```
type testCase struct {
    a, b float64
    want float64
}
```
and the use of **slices**, our new test functions will become as follows:<br/>
```
func TestXyz(t *testing.T) {
	testCases := []testCase{
		{a: A1, b: B1, want: W1},
		{a: A2, b: B2, want: W2},
		{a: A3, b: B3, want: W3},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Xyz(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Xyz(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
```
This will make it easier to create multiple test cases within each function.

First off, let's make it so that the tests will fail by configuring our `testCases` **wrong**.<br/>
Here's what the output would look like:<br/>
```
go test
--- FAIL: TestAdd (0.00s)
    calculator_test.go:24: Add(2.000000, 2.000000): want 5.000000, got 4.000000
    calculator_test.go:24: Add(1.000000, 1.000000): want 3.000000, got 2.000000
    calculator_test.go:24: Add(5.000000, 0.000000): want 6.000000, got 5.000000
--- FAIL: TestSubtract (0.00s)
    calculator_test.go:40: Subtract(101.000000, 100.000000): want 101.000000, got 1.000000
    calculator_test.go:40: Subtract(2.000000, 4.000000): want -1.000000, got -2.000000
    calculator_test.go:40: Subtract(10.000000, 3.000000): want 8.000000, got 7.000000
--- FAIL: TestMultiply (0.00s)
    calculator_test.go:56: Multiply(6.000000, 7.000000): want 43.000000, got 42.000000
    calculator_test.go:56: Multiply(1000.000000, 0.000000): want 1.000000, got 0.000000
    calculator_test.go:56: Multiply(3.000000, -3.000000): want -8.000000, got -9.000000
FAIL
exit status 1
FAIL	calculator	0.003s
```

