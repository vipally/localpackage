# projectroot
	this is a test of projectroot to allow local package reference style[import "#/xxx"]
	
	The aim is ./main works well anywhere projectroot is, even out of GoPath.
	
	./main <- #/local2 <- #/local1 <- #/public1
	./public2 <- github.com/vipally/localpackage/projectroot/public1