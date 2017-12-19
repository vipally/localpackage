package main

import (
	//cons "const/cons"
	"runtime"
	"time"
)

func main() {
	println("hello world!")
	println("buildtimestamp:", runtime.BuildTimestamp())
	println("buildtime     :", time.BuildTime().String())
	// Output:
	// hello world!
	// buildtimestamp: 1513671919
	// buildtime     : 2017-12-19 16:25:19 +0800 CST
}
