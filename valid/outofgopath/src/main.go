package main

import (
	"fmt"
	glocal "valid/localroot/src/locals/local1/local"

	"./local2"
)

func main() {
	glocal.Show()

	fmt.Println("main    import [./local2]")
	local2.PackageMain()
}
