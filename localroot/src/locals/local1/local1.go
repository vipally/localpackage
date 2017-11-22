package local1


import (
	"fmt"
	"#/publics/public1"
)

func PackageMain() {
	fmt.Println("local1  import [#/publics/public1]")
	public1.PackageMain()
}
