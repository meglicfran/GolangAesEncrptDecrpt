About GolangAesEncrptDecrpt
===========================
GolangAesEncrptDecrpt is a **windows** application that lets you encrypt and decrypt text using the AES algorithm in CBC mode with PKCS5Padding.

Setup
=====

Make sure you have Go and make installed.
See [Getting Started](http://golang.org/doc/install.html) and [Make for Windows](https://gnuwin32.sourceforge.net/packages/make.htm).

##### Build app

Install the [rsrc tool](https://github.com/akavel/rsrc):

	go install github.com/akavel/rsrc@latest

In the directory containing `main.go` run

	make

##### Run app

    AESEncryptDecrypt.exe