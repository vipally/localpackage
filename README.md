# localpackage
localpackage is a projcet to test local package reference problem.

***

# THE PROBLEM [The Discussion][gonuts]
	My project path is: github.com/vipally/localpackage/problem/withlocal
	The local package reference relation is:
	  withlocal <- withlocal/local

	In "main" use such to reference local package local:
	  import "github.com/VIPALLY/localpackage/problem/withlocal/local"
	
	Someone who forked this projcet as "github.com/someone/localpackage".
	But how can his project working by avoid following change?
	  import "import "github.com/SOMEONE/localpackage/problem/withlocal/local"

***

# My Solution [Reference][my_solution]
	Use [import "#/x/y/z"] local-reference style to related to LocalRoot instead of GoPath.
	
# Examples of with/without local package reference
	Here are two examples of "hello world" main packages:
	<ProjectRoot>/withoutlocal
	<ProjectRoot>/withlocal <- github.com/vipally/localpackage/local

	"withoutlocal" works well anywhere <ProjectRoot> is, even out of GoPath.
	"withlocal" works only when "<ProjectRoot> = <GoPath>/github.com/vipally/localpackage"
	How does go team think about this difference?
	
	It makes "withlocal" packages non-independent due to reference "LOCAL" packages as "GLOBAL" style.
	If I want my package works well anywhere, I have to write everything in one single "HUGE-package".
	Just like: all.Printf/all.OpenFile/all.GOROOT/...
	Does this go team recommended?
	
	We must explicit followed priorty of go dependency package find process:
	<ProjectRoot>: with highest-priorty path to find local packages.
	<Vendor>     : with second-priorty path to find explicit-version of local-referenced the third-party packages.
	<GoRoot>     : with third-priorty path to find standard packages.
	<GoPath>     : with lowest-priorty path to find third-party packages.
	
	Think about that not every go-project is wrote for open-source(aim to share with others).
	Thousands of private go-projects(eg:game-servers) focus on their own particular logic-flow only 
	and never shared private-packages at all.
	We just called these projects "independent-projects".
	Because they have hundreds-of private-packages but no one is wrote for share.

	That is to say, they never care "where I am", but "what I need".

	Unfortunately, these kind of projects are always "huge". 
	Maybe millions-of lines or thousands-of private-packages reference inside?

	In this case, change project name or source control server become heavy work, 
	because the working path changes and thousands-of private-packages reference code have to be update.

	But if local-packages are referenced by "#/modules/module1" style, 
	everything is change the name of project root only then.

	How do you think about the difference between such styles of referencing local-packages then?
	"#/modules/module1"
	"<GoRoot>/server/user/project/modules/module1"

[REFERENCE][examples]
	
## 1. Without local package reference
	"<ProjectRoot>/withoutlocal" is a main package without any local package reference
	It works well anywhere <ProjectRoot> is, even out of GoPath.

[REFERENCE][withoutlocal]

#### [ProjectRoot] = [GoPath]/github.com/vipally/localpackage
	Hello World! 
	WITHOUT local import
	Gopath is      : E:/gocode/src
	ProjectRoot is : E:/gocode/src/github.com/vipally/localpackage
	ThisPackagePath: E:/gocode/src/github.com/vipally/localpackage/withoutlocal
	ReleatGopath is: github.com/vipally/localpackage/withoutlocal

#### [ProjectRoot] = [GoPath]/../../localpackage
	Hello World! 
	WITHOUT local import
	Gopath is      : E:/gocode/src
	ProjectRoot is : E:/localpackage
	ThisPackagePath: E:/localpackage/withoutlocal
	ReleatGopath is: ../../localpackage/withoutlocal

***

## 2. With local package reference
	"<ProjectRoot>/withoutlocal" is a hello world package with local package reference:
	github.com/vipally/localpackage/local
	
	It works only when "<ProjectRoot> = <GoPath>/github.com/vipally/localpackage"

[REFERENCE][withlocal]

#### [ProjectRoot] = [GoPath]/github.com/vipally/localpackage
	Hello World! 
	WITH local import: github.com/vipally/localpackage/local
	Gopath is      : E:/gocode/src
	ProjectRoot is : E:/gocode/src/github.com/vipally/localpackage
	ThisPackagePath: E:/gocode/src/github.com/vipally/localpackage/withlocal
	ReleatGopath is: github.com/vipally/localpackage/withlocal

#### [ProjectRoot] = [GoPath]/../../localpackage
	withlocal.go:6:2: cannot find package "github.com/vipally/localpackage/local" in any of:
	C:\Go\src\github.com\vipally\localpackage\local (from $GOROOT)
	E:\gocode\src\github.com\vipally\localpackage\local (from $GOPATH)

***

# Maybe one solution [Reference][my_solution]
## 1.  use package comment to specify root of local project in projcetroot
	package main // import "#"


## 2. use someway to reference local package 
	import "#/local2"


[withlocal]: https://github.com/vipally/localpackage/blob/master/problem/withlocal/withlocal.go#L6
[withoutlocal]: https://github.com/vipally/localpackage/blob/master/problem/withoutlocal/withoutlocal.go#L9
[examples]: https://github.com/vipally/localpackage#examples-of-withwithout-local-package-reference
[gonuts]: https://groups.google.com/forum/#!topic/golang-nuts/ewyUOFyFIJU
[my_solution]: https://github.com/vipally/localpackage/tree/master/localroot/src#manage-private-only-projects-by-replacing-gopath-with-localroot
