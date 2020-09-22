## For The Love Of Go - Exercises
### Section 3: Fixing a hole

```
$ go test
--- FAIL: TestSubtract (0.00s)
    calculator_test.go:22: want 2.000000, got -2.000000
FAIL
exit status 1
FAIL	calculator	0.003s
```
Pay attention to `calculator_test.go:22` above, `22` is the row number we should be looking for.<br/>
Failure comes from the following section of `calculator_test.go`:
```
...
17 func TestSubtract(t *testing.T) {
18         t.Parallel()
19         var want float64 = 2
20         got := calculator.Subtract(4, 2)
21         if want != got {
22                 t.Errorf("want %f, got %f", want, got)   <<<
23         }
24 }
```
Evidently `calculator.Subtract(4, 2)` does not yield `2` as we'd expect. Let's take a look at the `Subtract()` function in `calculator.go`.</br>
```
func Subtract(a, b float64) float64 {
	return b - a
}
```
`Subtract(a, b)` implements the following operation, `b - a`. In our test, `a = 4`, `b = 2`, `b - a = 4 - 2 = -2`.</br>
To fix the code, we choose to reverse the operands as follows:</br>
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

**STRETCH GOAL**: Go provides an excellent documentation. Help can be displayed through CLI as follows:<br/>
```
$ go doc testing
package testing // import "testing"

type T struct {
	// Has unexported fields.
}
    T is a type passed to Test functions to manage test state and support
    formatted test logs.

    A test ends when its Test function returns or calls any of the methods
    FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods, as well as
    the Parallel method, must be called only from the goroutine running the Test
    function.

    The other reporting methods, such as the variations of Log and Error, may be
    called simultaneously from multiple goroutines.

func (c *T) Cleanup(f func())
func (t *T) Deadline() (deadline time.Time, ok bool)
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Helper()
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
func (c *T) TempDir() string
```

Once you're done with this section, continue with:</br>
`$ git checkout section_4`

