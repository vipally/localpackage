package main

import (
	//cons "const/cons"
	"runtime"
	"time"
)

func main() {
	println("hello world!")
	println("buildtime:", runtime.BuildTimestamp(), time.BuildTime().String())
	println("now time :", runtime.UnixNow(), time.Now().String())
	// Output:
	//hello world!
	//buildtime: 1513666403 2017-12-19 14:53:23 +0800 CST
	//now time : 1513666403 2017-12-19 14:53:23.9419813 +0800 CST m=+0.047002701
}
