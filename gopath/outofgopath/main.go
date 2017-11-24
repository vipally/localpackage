package main

import (
	"fmt"

	"./local"
)

func main() {
	println("Gopath/src")
	fmt.Println("main    import [local]")
	thisFilePath := local.GetThisFilepath(2)
	fmt.Printf("Gopath is      : %s\n", local.GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", local.GetProjectRootpath(2))
	fmt.Printf("ThisPackagePath: %s\n", local.GetThisFilepath(2))
	fmt.Printf("ReleatGopath is: %s\n", local.GetReleatGopath(thisFilePath))
}
