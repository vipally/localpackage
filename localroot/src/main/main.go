package main 

import (
	"fmt"
	 "#/util"
	"#/locals/local2"
)

func main() {
	TellImport("#/util")
	util.AssertRelatedLocalRoot("util")
	util.Show()

	TellImport("#/locals/local2")
	local2.AssertRelatedLocalRoot("locals/local2")
	local2.PackageMain()
}

func TellImport(imp string) {
	fmt.Printf("%s import [%s]\n", "main", imp)
}
