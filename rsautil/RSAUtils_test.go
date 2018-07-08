package rsautil

import (
	"testing"
	"encoding/base64"
)

func TestGenRsaKey(t *testing.T) {
	t.Log(GenRsaKey(2018))
}

func TestAll(t *testing.T) {
	pemPri, pemPub := GenRsaKey(1024)
	text := []byte("hello rsa")

	ciphertext, _ := Encrypt(text, pemPub)
	plainText, _ := Decrypt(ciphertext, pemPri)
	t.Log(string(plainText))

	sign, _ := Sign(text, pemPri)
	t.Log(base64.StdEncoding.EncodeToString(sign))

	err := VerifySign(text, sign, pemPub)
	t.Log(err)
}
