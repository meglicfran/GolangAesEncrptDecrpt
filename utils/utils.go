package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

type MyErr struct {
	msg string
}

func (err *MyErr) Error() string {
	return err.msg
}

func Sha256Hasing(input string) []byte {
	byteInput := []byte(input)
	sha256Hash := sha256.Sum256(byteInput)
	return sha256Hash[:]
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func AesEncrypt(input string, key string, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	if len(iv) != 16 {
		return nil, &MyErr{"Length of initialization vector must be 16 with AES."}
	}
	cbc := cipher.NewCBCEncrypter(block, iv)

	plaintext := []byte(input)
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	cbc.CryptBlocks(plaintext, plaintext)

	return plaintext, nil
}
