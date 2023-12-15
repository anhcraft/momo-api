package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func EncryptRSA(plainText []byte, publicKey []byte) ([]byte, error) {
	spkiBlock, _ := pem.Decode(publicKey)
	key, err := x509.ParsePKCS1PublicKey(spkiBlock.Bytes)
	if err != nil {
		return nil, err
	}
	out, err := rsa.EncryptPKCS1v15(rand.Reader, key, plainText)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func EncryptRSAToBase64(plainText []byte, publicKey []byte) (string, error) {
	data, err := EncryptRSA(plainText, publicKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func DecryptRSA(cipherText []byte, privateKey []byte) ([]byte, error) {
	spkiBlock, _ := pem.Decode(privateKey)
	key, err := x509.ParsePKCS1PrivateKey(spkiBlock.Bytes)
	if err != nil {
		return nil, err
	}
	out, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func DecryptRSAFromBase64(cipherTextEncoded string, privateKey []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherTextEncoded)
	if err != nil {
		return nil, err
	}
	data, err = DecryptRSA(data, privateKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}
