package main

import (
	"fmt"
	glocal "localroot1/src/util"

	"./local2"
)

func main() {
	glocal.Show()

	fmt.Println("main    import [./local2]")
	local2.PackageMain()
}
