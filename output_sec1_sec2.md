```
$ go test
PASS
ok  	calculator	0.002s
```

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

