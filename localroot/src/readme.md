# localroot
	This is a test of localroot refer local-only packages with [import "#/foo"] style.
	
- Content
	- [The Go patch]()
	- [What is a local-only package?]()
		- [1.With package comment [import "#"]]()
		- [2.With any import with [import "#/xxx"] style]()
	- [How to refer local-only packages]()
	- [What is LocalRoot?]()
	- [My test project. ]()
		- [1. LocalRoot under GoPath]()
		- [2. LocalRoot out of GoPath]()
		- [3. The local working tree]()

	
# Manage private-only projects by replacing GoPath with LocalRoot.
	
## The Go patch. [Code](https://github.com/vipally/go)
	I have made a patch of Golang to support [import "#/foo"] style reference.
	https://github.com/vipally/go
	
## What is a local-only package?
### 1.With package comment [import "#"] [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/locals/local1/local/local.go#L3)
	// [import "#"] declares this is a local-only package.
	// Though it doesnt have [import "#/foo"] style imports.
	package local //import "#"
### 2.With any import with [import "#/foo"] style  [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/locals/local2/local2.go#L5)
	import "#/locals/local1"

## How to compatible with current global import style? [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/witherrors/main.go#L4) 
	Local packages cannot referd by global style.(no matter from local and global packages)
	Global style import is with none limit in local packages.
	
## How to refer local-only packages? [Code](https://github.com/vipally/localpackage/blob/master/localroot/src/main/main.go#L6)
	With [import "#/x/y/z"] style.
	This style of import tells the compiler that it is a related package path based on LocalRoot, and never search GoPath.
	And it will be expanded as [import "x/y/z"] by compiler automatically related to LocalRoot but not GoPath.
	
## What is LocalRoot? [Code](https://github.com/vipally/go/blob/master/src/go/build/build.go#L393)
	LocalRoot is the root of a local-only project to replace "#" in [import "#/xxx"] style.
	LocalRoot is a <root> directory that contains such patten of sub-tree "<root>/src/vendor/" up from current path.
	A LocalRoot has the same tree structure with GoPath and GoRoot.

	Actually, a LocalRoot is a private GoPath that is accessible by sub-packages only.

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

## My test project. 
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
	
	Actually, a LocalRoot is a private GoPath that is accessible by sub-packages only.

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
	