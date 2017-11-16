# localpackage
localpackage is a projcet to test local package reference problem.

# The plobem
	My project path is: github.com/vipally/localpackage
	The local package reference relation is:
		main <- local2 <- local1

	In "main.go" use such to reference local package local2:
		import "github.com/VIPALLY/localpackage/local2"
	
	Someone who forked this projcet as "github.com/someone/localpackage".
	But how can his project working by avoid following change?
		import "github.com/SOMEONE/localpackage/local2"

### My project path
	github.com/vipally/localpackage

### This is the working tree of this project
	github.com/vipally/localpackage/
	│
	│  main.go
	│
	├─local1
	│      local1.go
	│
	└─local2
	       local2.go

	The local package reference relation is:
		main <- local2 <- local1
			
### Someone who forked my projcet
	github.com/someone/localpackage

	But how can his project working by avoid follow change?
		import "github.com/someone/localpackage/local2"
		
# Maybe one solution
## 1.  use package comment to specify root of local project in projcetroot
	package main // import "#"
	
## 2. use someway to reference local package 
	import "#/local2"
