package local1

import (
	"fmt"
)

import (
	"fmt"
	"#/public1"
)

func PackageMain() {
	fmt.Println("local2 import [#/public1]")
	public1.PackageMain()
}
