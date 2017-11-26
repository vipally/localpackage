package public2

import (
	"fmt"

	"valid/localroot/src/publics/public1"
)

func PackageMain() {
	fmt.Println("public2 import [valid/localroot/publics/public1]")
	public1.PackageMain()
}
