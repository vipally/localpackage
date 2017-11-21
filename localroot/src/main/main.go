package main //import "#"

import (
	//error: import local package by global style
	//"github.com/vipally/localpackage/projectroot/src/locals/local1/local"
	"#/locals/local1/local"
	"#/locals/local2"
	"fmt"
)

func main() {
	thisFilePath := local.GetThisFilepath(2)
	fmt.Printf("Gopath is      : %s\n", local.GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", local.GetProjectRootpath(2))
	fmt.Printf("ThisPackagePath: %s\n", local.GetThisFilepath(2))
	fmt.Printf("ReleatGopath is: %s\n", local.GetReleatGopath(thisFilePath))

	fmt.Println("*****************************")
	fmt.Println("main import [#/locals/local2]")
	local2.PackageMain()
	fmt.Println("*****************************")
}
