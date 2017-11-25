package public2

import (
	"fmt"

	"github.com/vipally/localpackage/projectroot/publics/public1"
)

func PackageMain() {
	fmt.Println("public2 import [github.com/vipally/localpackage/lprojectroot/publics/public1]")
	public1.PackageMain()
}
