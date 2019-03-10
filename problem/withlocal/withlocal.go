package main

import (
	"fmt"

	"problem/local"
)

// With local package import.
// It works only when <ProjectRoot> = <Gopath>/github.com/vipally/localpackage.
func main() {
	fmt.Println("Hello World! \nWITH local import:", local.GetReleatGopath(local.GetThisFilepath(1)))

	thisFilePath := local.GetThisFilepath(2)
	fmt.Printf("Gopath is      : %s\n", local.GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", local.GetProjectRootpath(2))
	fmt.Printf("ThisPackagePath: %s\n", local.GetThisFilepath(2))
	fmt.Printf("ReleatGopath is: %s\n", local.GetReleatGopath(thisFilePath))
}
