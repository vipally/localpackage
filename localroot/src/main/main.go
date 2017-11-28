package main //import "#"

import (
	"fmt"
	glocal "valid/localroot/src/locals/local1/local"
	"#/locals/local2"
)

func main() {
	fmt.Println("main    import [valid/localroot/src/locals/local1/local]")
	glocal.Show()

	fmt.Println("main    import [#/locals/local2]")
	local2.PackageMain()
}
