package local1

import (
	"fmt"
)

import (
	"fmt"
	"#/public2"
)

func PackageMain() {
	fmt.Println("local2 import [#/public2]")
	public2.PackageMain()
}
