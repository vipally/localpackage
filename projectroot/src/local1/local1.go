package local1


import (
	"fmt"
	"#/public1"
)

func PackageMain() {
	fmt.Println("local1 import [#/public1]")
	public1.PackageMain()
}
