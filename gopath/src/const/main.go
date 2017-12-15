package main

import (
	//cons "const/cons"
	"runtime"
	"time"
)

func main() {
	unix := runtime.BuildTimestamp()
	t := time.Unix(unix, 0)
	println("hello world")
	println("buildtimestamp:", unix)
	println("buildtime     :", t.String())
	// Output:
	// hello world
	// buildtimestamp: 1513329454
	// buildtime     : 2017-12-15 17:17:34 +0800 CST
}
