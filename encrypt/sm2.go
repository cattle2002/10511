package encrypt

import (
	"crypto/rand"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"os"
)

func Sm2PubEncrypt(key []byte) ([]byte, error) {
	file, err := os.ReadFile("sm2public.pem")
	if err != nil {
		return nil, err
	}
	fromPem, err := x509.ReadPublicKeyFromPem(file)
	if err != nil {
		return nil, err
	}
	cryptoRes, err := sm2.Encrypt(fromPem, key, rand.Reader, 1)
	if err != nil {
		return nil, err
	}
	return cryptoRes, nil
}
func Sm2PrivDecrypt(data []byte, key []byte) ([]byte, error) {
	file, err := os.ReadFile("sm2private.pem")
	if err != nil {
		return nil, err
	}

	fromPem, err := x509.ReadPrivateKeyFromPem(file, key)
	if err != nil {
		return nil, err
	}

	DcryptoRes, err := sm2.Decrypt(fromPem, data, 1)
	if err != nil {
		return nil, err
	}
	return DcryptoRes, err
}
