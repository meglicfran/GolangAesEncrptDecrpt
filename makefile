build: rsrc.syso test.go
	go build -ldflags="-H windowsgui"

rsrc.syso: 
	rsrc -manifest test.manifest -o rsrc.syso
