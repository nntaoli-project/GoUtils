package des

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = []byte("12345678")

func Test_DESEncrypt(t *testing.T) {
	data := []byte("hello")
	secretText, err := DESEncrypt(data, key)

	if err != nil {
		t.Error(err)
	}

	//output sxqYZRqwq+Q=
	assert.Equal(t, base64.StdEncoding.EncodeToString(secretText), "sxqYZRqwq+Q=")

}

func Test_DESDecrypt(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString("sxqYZRqwq+Q=")

	cleartext, err := DESDecrypt(data, key)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(cleartext), "hello")
}
