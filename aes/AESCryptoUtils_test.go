package aes

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = []byte("1234567812345678") // 16

func Test_DESEncrypt(t *testing.T) {
	data := []byte("hello")
	secretText, err := AESEncrypt(data, key, PKCS5Padding)

	if err != nil {
		t.Error(err)
	}

	//output vkYbld7ogr5nNi/bRp6XMA==
	assert.Equal(t, base64.StdEncoding.EncodeToString(secretText), "vkYbld7ogr5nNi/bRp6XMA==")

}

func Test_DESDecrypt(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString("vkYbld7ogr5nNi/bRp6XMA==")

	cleartext, err := AESDecrypt(data, key, nil)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(cleartext), "hello")
}
