build: rsrc.syso main.go
	go build -ldflags="-H windowsgui" -o AESEncryptDecrypt.exe

rsrc.syso: 
	rsrc -manifest test.manifest -o rsrc.syso
