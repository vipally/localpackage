package local2

import (
	"fmt"

	"github.com/vipally/localpackage/local1"
)

func Local2() {
	fmt.Println("local2 call local1:")
	local1.Local1()
}
