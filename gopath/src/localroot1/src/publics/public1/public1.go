package public1

import (
	"fmt"
	"go/build"
	localroot2 "localroot2/src/locals/local2"
	vendored2 "localrootvd1/src/locals/local2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"vendored" //"#/vendor/evndored"
)

var (
	gGoPath string
	gGoRoot string
)

func PackageMain() {
	fmt.Println("in package: " + GetPackagePath())
	TellImport("vendored (vendor tree)")
	AssertRelatedPath(GetLocalRoot(), "src/vendor/vendored", vendored.GetPackagePath())
	vendored.PackageMain()
	TellImport("localrootvd1/src/locals/local2 (vendor tree)")
	AssertRelatedPath(GetLocalRoot(), "localrootvd1/src/locals/local2", vendored.GetPackagePath())
	vendored2.PackageMain()
	TellImport("localroot2/src/locals/local2 (gopath)")
	AssertRelatedPath(GetLocalRoot(), "../localroot2/src/locals/local2", vendored.GetPackagePath())
	localroot2.PackageMain()

}

func GetPackageName() string {
	_, name := filepath.Split(GetPackagePath())
	return name
}

func TellImport(imp string) {
	fmt.Printf("%s import [%s]\n", GetPackageName(), imp)
}

func AssertRelatedGoPath(related string) {
	root := GetGopath()
	path := GetPackagePath()
	AssertRelatedPath(root, related, path)
}

func AssertRelatedGoRoot(related string) {
	root := GetGoRoot()
	path := GetPackagePath()
	AssertRelatedPath(root, related, path)
}

func AssertRelatedLocalRoot(related string) {
	root := filepath.Join(GetLocalRoot(), "src")
	path := GetPackagePath()
	AssertRelatedPath(root, related, path)
}

func AssertRelatedPath(root, related, dest string) {
	if filepath.Join(root, related) != filepath.Clean(dest) {
		panic(fmt.Sprintf("PathCheck: root=[%s][%s] dest=%s not match", root, related, dest))
	}
}

func GetPackagePath() string {
	return GetThisFilepath(1)
}

func GetGoRoot() string {
	if gGoRoot == "" {
		s := os.Getenv("GOROOT")
		if ss := strings.Split(s, ";"); ss != nil && len(ss) > 0 {
			gGoRoot = filepath.Clean(ss[0] + "/src")
		}
	}
	return gGoRoot
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

func GetLocalRoot() string {
	return GetLocalRootOfCaller(3)
}

func GetLocalRootOfCaller(depth int) string {
	return build.Default.SearchLocalRoot(filepath.Dir(GetThisFilepath(depth)))
}

func GetThisFilepath(depth int) string {
	thisFile, _ := FileLine(depth)
	thisFilePath := filepath.Dir(thisFile)
	return thisFilePath
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
