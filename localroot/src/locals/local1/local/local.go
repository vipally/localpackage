// import "#" declare this is a local package.
// Though it doesnt have "#/foo" style imports.
package local

//import "#"

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	gGoPath string
)

func Show() {
	fmt.Println("***********************************************************************************")
	fmt.Println("package local  : " + GetThisFilepath(1))
	fmt.Println("caller info    :")
	thisFilePath := GetThisFilepath(3)
	fmt.Printf("ThisPackagePath: %s\n", GetThisFilepath(3))
	fmt.Printf("ReleatGopath is: %s\n", GetReleatGopath(thisFilePath))
	fmt.Printf("Gopath is      : %s\n", GetGopath())
	fmt.Printf("ProjectRoot is : %s\n", GetProjectRootpath(4))
	fmt.Println("***********************************************************************************")
}

func GetGopath() string {
	if gGoPath == "" {
		s := os.Getenv("GOPATH")
		if ss := strings.Split(s, ";"); ss != nil && len(ss) > 0 {
			gGoPath = filepath.Clean(ss[0] + "/src")
		}
	}
	return gGoPath
}

func formatPath(s string) string {
	return filepath.ToSlash(filepath.Clean(s))
}

func GetProjectRootpath(depth int) string {
	return build.Default.SearchLocalRoot(filepath.Dir(GetThisFilepath(depth)))
}

func GetThisFilepath(depth int) string {
	thisFile, _ := FileLine(depth)
	thisFileath := filepath.Dir(thisFile)
	return thisFileath
}

func GetReleatGopath(s string) string {
	ss, _ := filepath.Rel(GetGopath(), s)
	return ss
}

//the caller's file/line info
func FileLine(depth int) (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(depth); __ok {
		file, line = __file, __line
	}
	return
}
