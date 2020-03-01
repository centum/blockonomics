package blockonomics

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func genPass(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func pad(data []byte) []byte {
	length := aes.BlockSize - len(data)%aes.BlockSize
	res := make([]byte, 0, len(data)+length)
	res = append(res, data...)
	c := byte(length)
	for i := 0; i < length; i++ {
		res = append(res, c)
	}
	return res
}

func unpad(data []byte) []byte {
	s := data[len(data)-1]
	return data[:s]
}

func bytesToKey(data, salt []byte, output int) []byte {
	if len(salt) != 8 {
		panic("salt length must be equal 8")
	}
	data = bytes.Join([][]byte{data, salt}, []byte{})
	key := md5.Sum(data)
	finalKey := append([]byte(nil), key[:]...)
	for {
		if len(finalKey) >= output {
			break
		}
		key = md5.Sum(bytes.Join([][]byte{key[:], data}, []byte{}))
		finalKey = bytes.Join([][]byte{finalKey, key[:]}, []byte{})
	}
	return finalKey[:output]
}

func Encrypt(plaintext []byte, passphrase []byte) []byte {
	salt := genPass(8)
	return saltingEncrypt(plaintext, passphrase, salt)
}

func saltingEncrypt(plaintext []byte, passphrase, salt []byte) []byte {
	keyIV := bytesToKey(passphrase, salt, 32+16)
	key := keyIV[:32]
	iv := keyIV[32:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	m := pad(plaintext)
	ciphertext := make([]byte, len(m))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, m)

	combineMsg := bytes.Join([][]byte{[]byte("Salted__"), salt, ciphertext}, []byte{})
	return []byte(base64.StdEncoding.EncodeToString(combineMsg))
}

func Decrypt(encrypted, passphrase []byte) []byte {
	encrypted, err := base64.StdEncoding.DecodeString(string(encrypted))
	if err != nil {
		panic(err)
	}
	if string(encrypted[0:8]) != "Salted__" {
		panic("wrong data")
	}
	salt := encrypted[8:16]
	keyIV := bytesToKey(passphrase, salt, 32+16)
	key := keyIV[:32]
	iv := keyIV[32:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(encrypted[16:]))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, encrypted[16:])
	return unpad(ciphertext)
}
