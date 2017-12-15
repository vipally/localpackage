package main

import (
	cons "const/cons"
	"runtime"
	"time"
)

func main() {
	println(cons.ConstVal)
	unix := runtime.BuildTimestamp()
	t := time.Unix(unix, 0)
	println(unix, t.String())
}
