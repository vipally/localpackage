//
// [import "#"] is someway to specify root of local project.
//
// package main // import "#"
//
//
package main

import (
	"fmt"

	// "./local2" //error: local import "./local2" in non-local package
)

// NOTE:
// This works well in my own procect
// But the problem is someone forked this project as "github.com/someone/localpackage"
// It doesn't work then.
// Because his working path is "github.com/someone/localpackage"
//
// [import "#/local2"] maybe a better way to reference local package.
// [#] will be replaced as proper working path who has been commented as [import "#"].
//
// import "#/local2"
//
import "github.com/vipally/localpackage/local2"

//working tree:
//	github.com/vipally/localpackage/
//	│
//	│  main.go
//	│
//	├─local1
//	│      local1.go
//	│
//	└─local2
//	       local2.go
//
//local package imort relation:
// 	main <- local2 <- local1

func main() {
	fmt.Println("main call local2:")
	local2.Local2()
}
