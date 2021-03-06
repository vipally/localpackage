package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	gGoPath string
)

func GetGopath() string {
	if gGoPath == "" {
		s := os.Getenv("GOPATH")
		if ss := strings.Split(s, ";"); ss != nil && len(ss) > 0 {
			gGoPath = formatPath(ss[0]) + "/src"
		}
	}
	return gGoPath
}

func formatPath(s string) string {
	return filepath.ToSlash(filepath.Clean(s))
}

func GetProjectRootpath(depth int) string {
	return formatPath(filepath.Dir(GetThisFilepath(depth)))
}

func GetThisFilepath(depth int) string {
	thisFile, _ := FileLine(depth)
	thisFilePath := filepath.Dir(thisFile)
	return formatPath(thisFilePath)
}

func GetReleatGopath(s string) string {
	ss, _ := filepath.Rel(GetGopath(), s)
	return formatPath(ss)
}

//the caller's file/line info
func FileLine(depth int) (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(depth); __ok {
		file, line = __file, __line
	}
	return
}
