package main

import (
	"fmt"
)

// Without local package import
// It works well anywhere of this package even out of Gopath.
func main() {
	fmt.Println("Hello World! \nWITHOUT local import")

	thisFilePath := GetThisFilepath(1)
	fmt.Printf("Gopath is      : %s\n", GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", GetProjectRootpath(1))
	fmt.Printf("ThisPackagePath: %s\n", GetThisFilepath(1))
	fmt.Printf("ReleatGopath is: %s\n", GetReleatGopath(thisFilePath))
}
