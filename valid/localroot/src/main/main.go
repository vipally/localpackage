package main //import "#"

import (
	//error: import local package by global style
	//"github.com/vipally/localpackage/localroot/src/locals/local1/local"
	"#/locals/local1/local"
	"#/locals/local2"
	"fmt"
	"golang.org/x/debug/macho"
)

func main() {
	local.Show()

	fmt.Println("**********************************************************")
	fmt.Println("main    import [#/locals/local2]")
	local2.PackageMain()
	fmt.Printf("main    import [golang.org/x/debug/macho(local vendor)]: \n%v\n",
		macho.CpuAmd64.GoString())
	fmt.Println("**********************************************************")
}
