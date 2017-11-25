package main //import "#"

import (
	//main.go:5:2: cannot import local-package E:\gocode\src\github.com\vipally\localpackage\localroot\src\locals\local2 from global style
	"github.com/vipally/localpackage/localroot/src/locals/local2"
	//"#/locals/local2"
)

func main() {
	local2.PackageMain()
}
