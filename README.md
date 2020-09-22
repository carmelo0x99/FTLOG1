## For The Love Of GO - Exercises
### 1: Testing times

Let's quickly try Go's testing features:</br>
```
$ go test
PASS
ok  	calculator	0.002s
```

Go also offers a tool to check your code is properly formatted:</br>
```
$ gofmt -d calculator.go
diff -u calculator.go.orig calculator.go
--- calculator.go.orig	2020-09-21 17:21:48.388830065 +0200
+++ calculator.go	2020-09-21 17:21:48.388830065 +0200
@@ -3,7 +3,7 @@
 
 // Add takes two numbers and returns the result of adding them together.
 func Add(a, b float64) float64 {
-return a + b
+	return a + b
 }
 
 // Subtract takes two numbers and returns the result of subtracting the second
```

To actually change in-place the code, run: `gofmt -w calculator.go`<br/>

Once you've familiarized with these features, continue by running:<br/>
`$ git checkout section_3`

