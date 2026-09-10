package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(value string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data := make([]byte, len(value)/2)

	_, err = fmt.Sscanf(value, "%x", &data)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize {
		return "", fmt.Errorf("invalid encrypted data")
	}

	nonce, encrypted := data[:nonceSize], data[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}