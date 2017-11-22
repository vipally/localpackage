# localroot
	this is a test of localroot to allow local package reference style[import "#/xxx"]
	
	The aim is ./main works well anywhere projectroot is, even out of GoPath.
	
	./main <- #/local2 <- #/local1 <- #/public1
	./public2 <- github.com/vipally/localpackage/projectroot/public1

#### 1. localroot under GoPath
	Gopath is      : E:\gocode\src
	ProjectRoot is : E:\gocode\src\github.com\vipally\localpackage\localroot
	ThisPackagePath: E:\gocode\src\github.com\vipally\localpackage\localroot\src\main
	ReleatGopath is: github.com\vipally\localpackage\localroot\src\main
	***********************************
	main import [#/locals/local2]
	local2 import [#/locals/local1]
	local1 import [#/publics/public1]
	public1(import noting)
	***********************************

#### 2. localroot out of GoPath
	Gopath is      : E:\gocode\src
	ProjectRoot is : E:\localpackage\localroot
	ThisPackagePath: E:\localpackage\localroot\src\main
	ReleatGopath is: ..\..\localpackage\localroot\src\main
	***********************************
	main import [#/locals/local2]
	local2 import [#/locals/local1]
	local1 import [#/publics/public1]
	public1(import noting)
	***********************************