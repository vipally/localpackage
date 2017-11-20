package public2

import (
	"fmt"

	"github.com/vipally/localpackage/projectroot/public1"
)

func PackageMain() {
	fmt.Println("public2 import [github.com/vipally/localpackage/lprojectroot/public1]")
	public1.PackageMain()
}
