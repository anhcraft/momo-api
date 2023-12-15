package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptAes256CbcPKCS7(plainText []byte, privateKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(privateKey)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	padding := aes.BlockSize - len(plainText)%aes.BlockSize
	if padding > 0 {
		paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
		plainText = append(plainText, paddingText...)
	}
	ciphertext := make([]byte, len(plainText))
	mode.CryptBlocks(ciphertext, plainText)
	return ciphertext, nil
}

func EncryptAes256CbcPKCS7ToBase64(plainText []byte, privateKey []byte) (string, error) {
	data, err := EncryptAes256CbcPKCS7(plainText, privateKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func DecryptAes256CbcPKCS7(cipherText []byte, privateKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(privateKey)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	lastByte := cipherText[len(cipherText)-1]
	padding := int(lastByte)
	if padding < aes.BlockSize {
		cipherText = cipherText[:len(cipherText)-padding]
	}
	return cipherText, nil
}

func DecryptAes256CbcPKCS7FromBase64(cipherTextEncoded string, privateKey []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherTextEncoded)
	if err != nil {
		return nil, err
	}
	data, err = DecryptAes256CbcPKCS7(data, privateKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}
