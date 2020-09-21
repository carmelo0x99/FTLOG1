Objective, add a `Multiply()` function. With TDD in mind we start off by creating the test first:<br/>
```
26 func TestMultiply(t *testing.T) {
27         t.Parallel()
28         var want float64 = 42
29         got := calculator.Multiply(6, 7)
30         if want != got {
31                 t.Errorf("want %f, got %f", want, got)
32         }
33 }
```

If we run `go test` now we'd, unsurprisingly, get an error:<br/>
```
# calculator_test [calculator.test]
./calculator_test.go:29:9: undefined: calculator.Multiply
FAIL	calculator [build failed]
```

The actual `Multiply` function must be written first:<br/>
```
16 func Multiply(a, b float64) float64 {
17 }
```

Again, running `go test` against this code will generate an error:<br/>
```
# calculator
./calculator.go:18:1: missing return at end of function
FAIL	calculator [build failed]
```
Let's now write a "good" function and run the test one more time:<br/>
```
16 func Multiply(a, b float64) float64 {
17         return a * b
18 }
```
and<br/>
```
PASS
ok  	calculator	0.003s
```

Very good! Now our function seems to be doing what we expected:
- `calculator_test.go` invokes `Multiply()` with the parameters `6` and `7`
- the expected outcome is `42`
- `PASS` is printed on screen<br/>

What would happen instead if we were to mess with the parameters or the expected resutl? Let's try...<br/>
```
28         var want float64 = 41
```
3, 2, 1...:<br/>
```
--- FAIL: TestMultiply (0.00s)
    calculator_test.go:31: want 41.000000, got 42.000000
FAIL
exit status 1
FAIL	calculator	0.002s
```
Good! Yes, that's actually good!!! What it says is that `Multiply(6, 7)` yields `42`. The number we've written in the test, `41`, doesn't match, an error is thrown.</br>
We can fix he test once and again and, wehn done, move to the next section with `git checkout section_5`.

