package crypto

import (
	"testing"
)

func TestAesEncrypt(t *testing.T) {

	t.Run("aes encrypt", func(t *testing.T) {
		var (
			// cipher key
			key = "dingtalkabcdefgh"
			// plaintext
			text = "This is a secret111asdfasdfsasdfasdfasdfasdfasdfsadf"
		)

		crypted, _ := EncryptByAes([]byte(text), key)
		str, _ := DecryptByAes(crypted, key)
		if text != string(str) {
			t.Error("aes encrypt error")
			return
		}
		t.Logf("crypted: %s;", crypted)
	})
}

func TestAesDecrypt(t *testing.T) {

}
