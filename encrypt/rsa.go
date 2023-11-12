package encrypt

import (
	"encoding/base64"
	"github.com/wumansgy/goEncrypt/rsa"
)

// 平台存放的Base64 公钥加密的密钥
func RsaDecrypt(CipherSymmetricKeyBase64 string, privateKey string) (string, error) {
	plainText, err := rsa.RsaDecryptByBase64(CipherSymmetricKeyBase64, privateKey)
	if err != nil {
		return "", err
	}
	toString := base64.StdEncoding.EncodeToString(plainText)
	return toString, nil
}
func RsaEncrypt() {}
