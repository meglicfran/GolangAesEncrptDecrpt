package utils

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// Copy/Pasta: got, gotErr := AesEncrypt([]byte("TestTestTest12345"), []byte("aesEncryptionKey"), []byte("1234567890123456"))
func TestAesEncrypt(t *testing.T) {
	//Invalid plaintext (nil)
	gotCiphertext, gotErr := AesEncrypt(nil, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invalid plaintext("")
	gotCiphertext, gotErr = AesEncrypt([]byte(""), []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invalid key(nil)
	gotCiphertext, gotErr = AesEncrypt([]byte("TestTestTest12345"), nil, []byte("1234567890123456"))
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invalid key (len!=16)
	gotCiphertext, gotErr = AesEncrypt([]byte("TestTestTest12345"), []byte("Key"), []byte("1234567890123456"))
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invalid IV (nil)
	gotCiphertext, gotErr = AesEncrypt([]byte("TestTestTest12345"), []byte("aesEncryptionKey"), nil)
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invalid IV (len != 16)
	gotCiphertext, gotErr = AesEncrypt([]byte("TestTestTest12345"), []byte("aesEncryptionKey"), []byte("123"))
	if gotErr == nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, gotErr, "<nil>, [error]")
	}

	//Invertibility Decrypt(Encrypt(x))==x
	ciphertext, err := AesEncrypt([]byte("TestTestTest12345"), []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		t.Errorf("Got %v,%v, want %v", ciphertext, err, "[ciphertext], <nil>")
	}
	gotPlaintext, err := AesDecrypt(ciphertext, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, err, "[plaintext], <nil>")
	}
	if !bytes.Equal(gotPlaintext, []byte("TestTestTest12345")) {
		t.Errorf("Got %v, want %v", gotPlaintext, "TestTestTest12345")
	}
}

//copy/pasta plaintext,err := AesDecrypt(ciphertext, []byte("aesEncryptionKey"), []byte("1234567890123456"))

func TestAesDecrypt(t *testing.T) {
	ciphertext, err := hex.DecodeString("2fc574b6921f123abadaf4db734e7462c259a33abedb08f5e42b0697c3f0c17a")
	if err != nil {
		t.Error(err)
		return
	}

	//invalid cipherText (nil)
	plaintext, err := AesDecrypt(nil, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid cipherText ("")
	plaintext, err = AesDecrypt([]byte(""), []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid cipherText (len != k*16)
	plaintext, err = AesDecrypt([]byte("43234r234r234r24radf"), []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid key (nil)
	plaintext, err = AesDecrypt(ciphertext, nil, []byte("1234567890123456"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid key (len!=16)
	plaintext, err = AesDecrypt(ciphertext, []byte("Key"), []byte("1234567890123456"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid IV (nil)
	plaintext, err = AesDecrypt(ciphertext, []byte("Key"), nil)
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//invalid IV (len != 16)
	plaintext, err = AesDecrypt(ciphertext, []byte("Key"), []byte("123"))
	if err == nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "<nil>, [error]")
	}

	//Invertibility Encrypt(Decrypt(x))==x
	plaintext, err = AesDecrypt(ciphertext, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		t.Errorf("Got %v,%v, want %v", plaintext, err, "[plaintext], <nil>")
	}
	gotCiphertext, err := AesEncrypt(plaintext, []byte("aesEncryptionKey"), []byte("1234567890123456"))
	if err != nil {
		t.Errorf("Got %v,%v, want %v", gotCiphertext, err, "[ciphertext], <nil>")
	}
	if !bytes.Equal(ciphertext, gotCiphertext) {
		t.Errorf("Got %s, want %s", gotCiphertext, ciphertext)
	}

}
