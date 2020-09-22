Let's now improve upon our test strategy by making the tests more modular. We can achieve that quite easily by introducing a `struct`, such as:<br/>
```
 8 type testCase struct {
 9         a, b float64
10         want float64
11 }
```
... and by using **slices** to define our rewritten test functions as follows:<br/>
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
Overall, this will make it easier to create multiple test cases within each function.

First off, let's make it so that the tests will fail by configuring our `testCases` with the **wrong** values for `want`.<br/>
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
The information that gets displayed is very useful. We can actually see the parameters we're passing to the functions, its output and the expected result.<br/>
Next, we can fix the values for `want` and check whether the test is passed:<br/>
```
PASS
ok  	calculator	0.003s
```
To achieve the stretch goal of adding a `name` to our tests we must modify the following:<br/>
- add "fmt" to our import list
```
 3 import (
 4         "calculator"
 5         "fmt"           <<<
 6         "testing"
 7 )
```
- add an appropriate field to the struct's definition:
```
 9 type testCase struct {
10         a, b float64
11         want float64
12         name string     <<<
13 }
```
- expand each test function with additional code checking for the existance of a field named `name` and print it:<br/>
```
 25         if tc.want != got {
 26                 if tc.name != "" {                                     <<<
 27                         fmt.Println("[!] Failed test: ", tc.name)      <<<
 28                 }                                                      <<<
 29                 t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
 30         }
```
With the changes above, errors would be reported along with their `names`:<br/>
```
[!] Failed test:  2 + 2 = 4
[!] Failed test:  Douglas was here
[!] Failed test:  101 - 100 = 1
--- FAIL: TestAdd (0.00s)
    calculator_test.go:29: Add(2.000000, 2.000000): want 3.000000, got 4.000000
--- FAIL: TestMultiply (0.00s)
    calculator_test.go:67: Multiply(6.000000, 7.000000): want 43.000000, got 42.000000
--- FAIL: TestSubtract (0.00s)
    calculator_test.go:48: Subtract(101.000000, 100.000000): want 111.000000, got 1.000000
FAIL
exit status 1
FAIL	calculator	0.003s
```
**Cryptic note**: `t.Parallel()` is not our friend here<br/>
The next section can be reached through:<br/>
`$ git checkout section_6`

