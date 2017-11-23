# localroot
	This is a test of localroot refer local-only packages with [import "#/foo"] style.

# Manage private-only golang-projects by replacing GoPath with LocalRoot.
	
- Content
	- [The problem](https://github.com/vipally/localpackage/tree/master/localroot/src#the-problem-refer)
	- [The Go patch](https://github.com/vipally/localpackage/tree/master/localroot/src#the-go-patch-code)
	- [The solution](https://github.com/vipally/localpackage/tree/master/localroot/src#the-solution)
	- [What is a local-only package?](https://github.com/vipally/localpackage/tree/master/localroot/src#what-is-a-local-only-package)
		- [1.With package comment [import "#"]](https://github.com/vipally/localpackage/tree/master/localroot/src#1with-package-comment-import--code)
		- [2.With any import with [import "#/xxx"] style](https://github.com/vipally/localpackage/tree/master/localroot/src#2with-any-import-with-import-foo-style--code)
	- [How to compatible with current global import style?](https://github.com/vipally/localpackage/tree/master/localroot/src#how-to-compatible-with-current-global-import-style-code)
	- [How to refer local-only packages](https://github.com/vipally/localpackage/tree/master/localroot/src#how-to-refer-local-only-packages-code)
	- [What is LocalRoot?](https://github.com/vipally/localpackage/tree/master/localroot/src#what-is-localroot-code)
	- [My test project. ](https://github.com/vipally/localpackage/tree/master/localroot/src#my-test-project)
		- [1. LocalRoot under GoPath](https://github.com/vipally/localpackage/tree/master/localroot/src#1-localroot-under-gopath)
		- [2. LocalRoot out of GoPath](https://github.com/vipally/localpackage/tree/master/localroot/src#2-localroot-out-of-gopath)
		- [3. The local working tree](https://github.com/vipally/localpackage/tree/master/localroot/src#3-the-local-working-tree)

## The problem. [Refer](https://github.com/vipally/localpackage#the-problem)

## The Go patch. [Code](https://github.com/vipally/go) [TestCode][localroot]
	I have made a patch of Golang to allow refering local packages by [import "#/foo"] style.
	https://github.com/vipally/go

## The solution
### 1.  use such way to define a [LocalRoot][SearchLocalRoot] to replacing GoPath/GoRoot [Examle][localroot]
	LocalRoot is a <root> directory that contains such patten of sub-tree "<root>/src/vendor/" up from current path.
	A LocalRoot has the same tree structure with GoPath and GoRoot.

	Actually, a LocalRoot is a private GoPath that is accessible to sub-packages only.

	This is the minimal state of a valid LocalRoot:
		LocalRoot
		│
		└─src
		    ├─vendor
		    │  ...
		    └─...

	After build and install, it may become as:
		LocalRoot
		│  
		├─bin
		│    ...
		├─pkg
		│  └─windows_amd64
		│      └─...
		└─src
		    │  ...     
		    ├─vendor
		    │  └─...
		    └─...

### 2. use [import "#/x/y/z"] style to refer local package [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/locals/local2/local2.go#L5)
	import "#/x/y/z"
	
	Which means [import "x/y/z"] from LocalRoot, and never search from GoPath/GoRoot.
	
## What is a local-only package?
### 1.With package comment [import "#"] [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/locals/local1/local/local.go#L3)
	// [import "#"] declares this is a local-only package.
	// Though it doesnt have [import "#/foo"] style imports.
	package local //import "#"
### 2.With any import with [import "#/foo"] style  [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/locals/local2/local2.go#L5)
	import "#/locals/local1"

## How to compatible with current global import style? [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/witherrors/main.go#L4) 
	Local packages cannot refered by global style(no matter by local and global packages).
	Global style import is with none limitation in local packages.
	Source tree that cannot find a pattern of "<root>/src/vendor/" sub-tree cannot use local import style.
	
## How to refer local-only packages? [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/main/main.go#L6)
	With [import "#/x/y/z"] style.
	
	Which means [import "x/y/z"] from LocalRoot, and never search from GoPath/GoRoot.
	This style of import tells the compiler that it is a local package related to LocalRoot, 
	and never search from GoPath/GoRoot.
	It will be expanded as [import "x/y/z"] by compiler automatically related to LocalRoot but not GoPath/GoRoot.
	
## What is LocalRoot? [Code][SearchLocalRoot]
	LocalRoot is the root of a local-only project to replace "#" in [import "#/xxx"] style.
	LocalRoot is a <root> directory that contains such patten of sub-tree "<root>/src/vendor/" up from current path.
	A LocalRoot has the same tree structure with GoPath and GoRoot.

	Actually, a LocalRoot is a private GoPath that is accessible to sub-packages only.

	This is the minimal state of a valid LocalRoot:
		LocalRoot
		│
		└─src
		    ├─vendor
		    │  ...
		    └─...
	
	After build and install, it may become as:
		LocalRoot
		│  
		├─bin
		│    ...
		├─pkg
		│  └─windows_amd64
		│      └─...
		└─src
		    │  ...     
		    ├─vendor
		    │  └─...
		    └─...

## My test project. [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/main/main.go#L1)
	As my expectation, it works well wherever LocalRoot is, even out of GoPath.
	Actually, a LocalRoot is a private GoPath that is accessible to sub-packages only.

	If a private project use local and vendor third-party packages only,
	it will have nothing to do with GoPath.

### 1. LocalRoot under GoPath
	Gopath is      : E:\gocode\src
	ProjectRoot is : E:\gocode\src\github.com\vipally\localpackage\localroot
	ThisPackagePath: E:\gocode\src\github.com\vipally\localpackage\localroot\src\main
	ReleatGopath is: github.com\vipally\localpackage\localroot\src\main
	**********************************************************
	main    import [#/locals/local2]
	local2  import [#/locals/local1]
	local1  import [#/publics/public1]
	public1 import noting
	main    import [golang.org/x/debug/macho(local vendor)]: 
	macho.CpuAmd64
	**********************************************************

### 2. LocalRoot out of GoPath
	Gopath is      : E:\gocode\src
	ProjectRoot is : E:\localpackage\localroot
	ThisPackagePath: E:\localpackage\localroot\src\main
	ReleatGopath is: ..\..\localpackage\localroot\src\main
	**********************************************************
	main    import [#/locals/local2]
	local2  import [#/locals/local1]
	local1  import [#/publics/public1]
	public1 import noting
	main    import [golang.org/x/debug/macho(local vendor)]: 
	macho.CpuAmd64
	**********************************************************
	
### 3. The local working tree
	This is the tree of LocalRoot after build/install.
	
	Actually, a LocalRoot is a private GoPath that is accessible to sub-packages only.

	LocalRoot
	│
	├─bin
	│     main.exe
	├─pkg
	│  └─windows_amd64
	│      ├─locals
	│      │  └─local1
	│      ├─publics
	│      └─vendor
	│          └─golang.org
	│              └─x
	│                  └─debug
	└─src
	    ├─locals
	    │  ├─local1
	    │  │  └─local
	    │  └─local2
	    ├─main
	    ├─publics
	    │  ├─public1
	    │  └─public2
	    └─vendor
	      └─golang.org
	          └─x
	              └─debug
	                  ├─macho
	                  └─...
	
[BackToTop](https://github.com/vipally/localpackage/tree/master/localroot/src#localroot)

[SearchLocalRoot]: https://github.com/vipally/go/blob/master/src/go/build/build.go#L397
[localroot]: https://github.com/vipally/localpackage/tree/master/localroot/src