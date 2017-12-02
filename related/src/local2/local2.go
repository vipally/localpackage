package local2

//import "#"

import (
	"fmt"
	"path/filepath"
	"runtime"

	"../local1"
)

func PackageMain() {
	fmt.Println("in package: " + getThisFilepath(1))
	fmt.Println("import [../local1]")
	local1.PackageMain()
}

func getThisFilepath(depth int) string {
	thisFile, _ := fileLine(depth)
	thisFilePath := filepath.Dir(thisFile)
	return thisFilePath
}

//the caller's file/line info
func fileLine(depth int) (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(depth); __ok {
		file, line = __file, __line
	}
	return
}
