package main

import (
	//"github.com/vipally/localpackage/lprojectroot/public2"
	"#/local2"
	"#/local1/local"
	"fmt"
)

func main() {
	thisFilePath := local.GetThisFilepath(2)
	fmt.Printf("Gopath is      : %s\n", local.GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", local.GetProjectRootpath(2))
	fmt.Printf("ThisPackagePath: %s\n", local.GetThisFilepath(2))
	fmt.Printf("ReleatGopath is: %s\n", local.GetReleatGopath(thisFilePath))

	fmt.Println("\n*************************")
	fmt.Println("main import [#/local2]")
	local2.PackageMain()
//	fmt.Println("local2 import [github.com/vipally/localpackage/lprojectroot/public2]")
//	public2.PackageMain()
	fmt.Println("*************************")
}
