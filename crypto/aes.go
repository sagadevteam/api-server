package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// Encrypt the message with key
func Encrypt(key []byte, msg string) (enc string, err error) {
	pt := []byte(msg)

	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	ct := make([]byte, aes.BlockSize+len(pt))
	iv := ct[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(b, iv)
	stream.XORKeyStream(ct[aes.BlockSize:], pt)
	enc = hex.EncodeToString(ct)
	return
}

// Decrypt the message with key
func Decrypt(key []byte, enc string) (msg string, err error) {
	ct, err := hex.DecodeString(enc)
	if err != nil {
		return
	}

	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(ct) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short")
		return
	}

	iv := ct[:aes.BlockSize]
	ct = ct[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(b, iv)
	stream.XORKeyStream(ct, ct)
	msg = string(ct)
	return
}
