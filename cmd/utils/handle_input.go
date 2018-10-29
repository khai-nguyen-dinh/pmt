package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func hashTo32Bytes(input string) (output string) {

	if len(input) == 0 {
		return ""
	}

	hasher := sha256.New()
	hasher.Write([]byte(input))

	stringToSHA256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// Cut the length down to 32 bytes and return.
	return stringToSHA256[:32]
}

func DecryptString(cryptoText string, keyString string) (plainTextString string, err error) {

	encrypted, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}
	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short. It decodes to %v bytes but the minimum length is 16", len(encrypted))
	}

	decrypted, err := decryptAES([]byte(hashTo32Bytes(keyString)), encrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func decryptAES(key, data []byte) ([]byte, error) {
	// split the input up in to the IV seed and then the actual encrypted data.
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(data, data)
	return data, nil
}

func EncryptString(plainText string, keyString string) (cipherTextString string, err error) {

	key := hashTo32Bytes(keyString)
	encrypted, err := encryptAES([]byte(key), []byte(plainText))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func encryptAES(key, data []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	output := make([]byte, aes.BlockSize+len(data))
	iv := output[:aes.BlockSize]
	encrypted := output[aes.BlockSize:]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(encrypted, data)
	return output, nil
}
