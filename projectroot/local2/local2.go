package local2

import (
	"fmt"
	"#/local1"
)

func PackageMain() {
	fmt.Println("local2 import [#/local1]")
	local1.PackageMain()
}
