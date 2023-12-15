package main

import (
	"github.com/stretchr/testify/assert"
	"momo-api/utils/crypto"
	"testing"
)

func TestAES(t *testing.T) {
	key := []byte("Z3jk7Y+Ya4yjh3cVpfG2mSqD7QjvBtfCMOQW59u/PlaXuGfPbXNN9rWc5sLPzYwO"[:32])
	n, err := crypto.EncryptAes256CbcPKCS7([]byte("abcd"), key)
	if err != nil {
		t.Error(err)
	}
	m, f := crypto.DecryptAes256CbcPKCS7(n, key)
	if f != nil {
		t.Error(f)
	}
	assert.Equal(t, "abcd", string(m))
}

func TestRSA(t *testing.T) {
	publicKey := []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZBTz0r0JdxVXZuf2N8wWBahjl\nwlNJIpcjN8uTARd/2klOhOAIfiOtP1i5eKS5okSK/zmhyYeT9je66MfOBv3gdhIe\noO4TIOwitFShn7e97i1mZS7vzimwt6yWR69R2AGPXPKPA1AGdXo7kfk4CgvA0+ru\nUNTt/eMNQYDmW8W0pQIDAQAB\n-----END PUBLIC KEY-----")
	privateKey := []byte("-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQCZBTz0r0JdxVXZuf2N8wWBahjlwlNJIpcjN8uTARd/2klOhOAI\nfiOtP1i5eKS5okSK/zmhyYeT9je66MfOBv3gdhIeoO4TIOwitFShn7e97i1mZS7v\nzimwt6yWR69R2AGPXPKPA1AGdXo7kfk4CgvA0+ruUNTt/eMNQYDmW8W0pQIDAQAB\nAoGBAJXvCv4Zp+anj+opkqb+43sd0U13bhHwIbUxW3gsDrGI2mFkUvwfVKfRtQbu\njkESHSUw1XCQPhcdnxe2NOjL8v7zhFwkrcvjZsXWADhxgcfN8o+ZWq3BlRt7UZpH\nUPnfVjMe8GupeyWeKmP2moZbU5jyOOV6JUjvhMgh6J+Ut6kBAkEAz4vlEJ+5+aqu\naUqg8VDrCeBWkPJRPSGoMKbfU171DB9TYfnBSCNbsPfh9OTUjaDm6nLocYVSwjzz\nw1j9UV0NwQJBALy+kqsPdxmWQqsyIUz++K+Yu8KC3/xX6qyMRKyxiQjYvYhXwhxP\nFkf7xQddWQF8jf3jjQi5i9FMmgStTeJwJ+UCQErpfLGmZXMnVzKr/DF9+ogjEDvb\nKtV8239MDBnEkYBhojAf/NKz6HmUz1scaVgBdrey6BFphPiVFYsyCKUgiEECQQCT\n7HwAiv9Z01Tu3TwSHyaCYJ6O5IltOO4YS1qrSfzPLSbmC3l7PFSHGAAkNHnEW3zh\nRYzMELdO0s1G6xhGZoYtAkEAtnfuxPpTeD6HW9cMUNJqGzFkWxpnx02EYU0tKvdt\n48D4KTaGnMx1qHo/oiUUaEbcXY/XIWltW+l60NMoVDi6Rg==\n-----END RSA PRIVATE KEY-----")
	n, e := crypto.EncryptRSA([]byte("abcd"), publicKey)
	if e != nil {
		t.Error(e)
	}
	m, f := crypto.DecryptRSA(n, privateKey)
	if f != nil {
		t.Error(f)
	}
	assert.Equal(t, "abcd", string(m))
}
