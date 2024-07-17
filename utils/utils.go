package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
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

func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length <= 0 {
		return nil, fmt.Errorf("invalid byte blob lenght: expecting > 0 having %d", length)
	}
	unpadding := int(src[length-1])
	delta := length - unpadding
	if delta < 0 {
		return nil, fmt.Errorf("invalid padding delta lenght: expecting >= 0 having %d", delta)
	}
	return src[:delta], nil
}

// AesEncrypt returns the []byte representing the ciphertext of the plaintext
// encrypted with AES|CBC|PKCS5Padding using key and iv.
func AesEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	if plaintext == nil || len(plaintext) < 1 {
		return nil, &MyErr{"plaintext can't be nil or empty"}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != 16 {
		return nil, &MyErr{"length of initialization vector must be 16 with AES."}
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	cbc.CryptBlocks(plaintext, plaintext)

	return plaintext, nil
}

// AesDecrypt decrypts cipertext []byte encrypted with AES|CBC|PKCS5Padding using key and iv.
// Returns plaintext []byte or error.
func AesDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	if ciphertext == nil || len(ciphertext)%aes.BlockSize != 0 {
		return nil, &MyErr{"cyphertext size must be a multiple of block size"}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != aes.BlockSize {
		return nil, &MyErr{"length of initialization vector must be 16 with AES."}
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, ciphertext)
	plaintext, err := PKCS5UnPadding(ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
