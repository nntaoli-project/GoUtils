package rsautil

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"crypto/sha256"
	"crypto"
)

type ByteBuf struct {
	data []byte
}

func (w *ByteBuf) Write(p []byte) (n int, err error) {
	w.data = append(w.data, p...)
	return 1, nil
}

func GenRsaKey(bit int) (string, string) {
	pri, _ := rsa.GenerateKey(rand.Reader, bit)
	derStream := x509.MarshalPKCS1PrivateKey(pri)

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	buf := new(ByteBuf)

	pem.Encode(buf, block)

	pemPri := string(buf.data)

	pubKey := &pri.PublicKey
	pubPkix := x509.MarshalPKCS1PublicKey(pubKey)
	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubPkix,
	}

	buf2 := new(ByteBuf)
	pem.Encode(buf2, pubBlock)

	pemPub := string(buf2.data)
	return pemPri, pemPub
}

// 加密
func Encrypt(origData []byte, pemKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, origData, []byte(""))
}

// 解密
func Decrypt(ciphertext []byte, pemKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, []byte(""))
}

func Sign(text []byte, pemKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

	hash := crypto.SHA256
	hash2 := hash.New()
	hash2.Write(text)
	hashed := hash2.Sum(nil)

	return rsa.SignPSS(rand.Reader, priv, hash, hashed, &opts)
}

func VerifySign(text []byte, sign []byte, pemKey string) error {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return errors.New("public key error!")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example

	hash := crypto.SHA256
	hash2 := hash.New()
	hash2.Write(text)
	hashed := hash2.Sum(nil)

	return rsa.VerifyPSS(pubKey, hash, hashed, sign, &opts)
}
