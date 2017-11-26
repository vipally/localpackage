package main //import "#"

import (
	"fmt"
	glocal "valid/localroot/src/locals/local1/local"
	"#/locals/local2"

	"golang.org/x/debug/macho"
)

func main() {
	glocal.Show()

	fmt.Println("**********************************************************")
	fmt.Println("main    import [#/locals/local2]")
	local2.PackageMain()
	fmt.Printf("main    import [golang.org/x/debug/macho(local vendor)]: \n%v\n",
		macho.CpuAmd64.GoString())
	fmt.Println("**********************************************************")
}
