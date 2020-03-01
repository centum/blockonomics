package blockonomics

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Run("encrypt", func(t *testing.T) {
		pass := genPass(8)
		message := "Hi Blockonomics!" //nolint:goconst
		r := Encrypt([]byte(message), pass)
		if len(r) == 0 {
			t.Error("result empty")
		}
	})

	t.Run("salt encrypt", func(t *testing.T) {
		pass := "ONRhfKsU"
		salt := "23433KsU"
		message := "Hi Blockonomics!" //nolint:goconst

		expect := "553246736447566b583138794d7a517a4d30747a56653144726a2f784c6d6e48684b496a634a626e382b426456535630543032515474786a325a343145385058" //nolint:lll
		got := saltingEncrypt([]byte(message), []byte(pass), []byte(salt))
		if expect != fmt.Sprintf("%x", got) {
			t.Errorf("\nexpect: %s\n   got: %x\n", expect, got)
		}

	})
	t.Run("bytes to key", func(t *testing.T) {
		key := "12345678"
		salt := "SALTDFGH"
		got := bytesToKey([]byte(key), []byte(salt), 32+16)
		expect := "7dc4b7d90c2494b34629b895dc68f70c53dc9d13b59dd1d8271e9bb3aacec136ad4b14680d30824a6fa4845d9f1bfdf5"
		if expect != fmt.Sprintf("%x", got) {
			t.Errorf("\nexpect: %s\n   got: %x\n", expect, got)
		}

	})
	t.Run("padding", func(t *testing.T) {
		message := "Hi Blockonomics!" //nolint:goconst
		expect := "Hi Blockonomics!"
		got := pad([]byte(message))
		if expect != string(got) {
			t.Errorf("\nexpect: %s\n   got: %s\n", expect, got)
		}

	})
	t.Run("encrypt -> decrypt", func(t *testing.T) {
		pass := genPass(8)
		message := "Hi Blockonomics!" //nolint:goconst
		encrypted := Encrypt([]byte(message), pass)
		if len(encrypted) == 0 {
			t.Error("result empty")
		}

		got := Decrypt(encrypted, pass)

		if message != string(got) {
			t.Errorf("\nexpect: %s\n   got: %s\n", message, got)
		}
	})
}
