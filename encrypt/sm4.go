package encrypt

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"github.com/tjfoc/gmsm/sm4"
)

func Sm4Decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个block模式
	blockMode := cipher.NewCBCDecrypter(block, make([]byte, sm4.BlockSize))
	//decodeString, _ := base64.StdEncoding.DecodeString(ciphertext)
	// 解密
	decryptedText := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(decryptedText, ciphertext)

	// 去除填充
	decryptedText, err = PKCS7UnPadding(decryptedText, sm4.BlockSize)

	return decryptedText, err
}
func PKCS7UnPadding(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, errors.New("invalid plaintext")
	}
	return plainText[:length-number], nil
}

func PKCS7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}
