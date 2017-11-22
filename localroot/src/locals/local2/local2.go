package local2

import (
	"fmt"
	"#/locals/local1"
)

func PackageMain() {
	fmt.Println("local2  import [#/locals/local1]")
	local1.PackageMain()
}
