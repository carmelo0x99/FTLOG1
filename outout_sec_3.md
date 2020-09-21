### Section 3: Fixing a hole
```
$ go test
--- FAIL: TestSubtract (0.00s)
    calculator_test.go:22: want 2.000000, got -2.000000
FAIL
exit status 1
FAIL	calculator	0.003s
```

Failure comes from the following section of `calculator_test.go`:
```
...
17 func TestSubtract(t *testing.T) {
18         t.Parallel()
19         var want float64 = 2
20         got := calculator.Subtract(4, 2)
21         if want != got {
22                 t.Errorf("want %f, got %f", want, got)
23         }
24 }
```
Evidently `calculator.Subtract(4, 2)` does not yield `2` as we'd expect. Let's take a look at the `Subtract()` function in `calculator.go`.</br>
```
func Subtract(a, b float64) float64 {
	return b - a
}
```
`Subtract(a, b)` implements the following operation, `b - a`. In out test, `a = 4`, `b = 2`, `b - a = 4 - 2 = -2`.</br>
To fix the code, we may reverse the operands as follows:</br>
```
func Subtract(a, b float64) float64 {
        return a - b
}
```
Once that's done, we can run the test again and the result is:</br>
```
PASS
ok  	calculator	0.003s
```

