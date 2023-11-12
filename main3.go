package main

import (
	"DIMSMonitorPlat/utils"
	"fmt"
	"github.com/cattle2002/easycrypto/ecrypto"
)

func main33() {
	//download, err := oftp.Download("aes田馥甄小幸运kgma.txt")
	//download, err := oftp.Download("admin-1697699950428-d48ae022b7d679f19ecde2e4b60422af-47-468.enc")
	//
	//fmt.Println(err, download)
	//file, err := os.ReadFile("product.txt")
	//fmt.Println("2", err)
	////
	//decrypt, err := ecrypto.SymmetricDecrypt("aes", file)
	//fmt.Println("1", err)
	//create, _ := os.Create("test.kgma")
	//reader := bytes.NewReader(decrypt)
	//written, err := io.Copy(create, reader)
	//fmt.Println("3", written, err)
	println(utils.GetTime())
	//rsa.GenerateKey()
	//ecrypto.RsaPublicKeyEncryptSymmetricKey()
	//rsa.GenerateRsaKeyBase64()
	pk := "MIGJAoGBAM3biJ5YMY7typdu2btxdhnpU45l8D59XT+0XzhQ3Z4CkJGrWemsMTyu6OkbdWi8Czk7+cPgpIaYUuZ6Eb2hUjYRt/18OFzrw8WV4KeIbZsCzAkCmVvO8/sCVm6K7HfVMWqPu9DRw5qoBoG+s/mahP6BD+NfA4G3WOFc+hWp6sndAgMBAAE="
	sk := "MIICXQIBAAKBgQDN24ieWDGO7cqXbtm7cXYZ6VOOZfA+fV0/tF84UN2eApCRq1nprDE8rujpG3VovAs5O/nD4KSGmFLmehG9oVI2Ebf9fDhc68PFleCniG2bAswJAplbzvP7AlZuiux31TFqj7vQ0cOaqAaBvrP5moT+gQ/jXwOBt1jhXPoVqerJ3QIDAQABAoGARwEa/Wr0JRWWFGeSyFmZX9SgEnTMyfEZ0dhYI6eY2WCPFv+gcr4G+aRnB8iF1ctOn741Jz/738T4izC9n8fd/9O0MsBJYCFHg2y8vlEaUxqZj/eRz0Z4L4OmGoXR3nPq0syFptA3wSg4n2yybUpYCXUx4hqeOZOv/bKOpb4HYm0CQQDYLLCwbeE+J1f5oU6XSBR1l/AAAYXXu4weDI72MLq0Ulkt40QHnSlGabMl2geLZej5fBn2NYBRBq1bO9rA5Vv3AkEA88hEq6wJeF2CnYCrQ75L6XXMawuijCBvkrgtHMCsYZn9MoCqKwF/KfyqTAwgXooYeLFGAwvnwtP4nNhr2VLLywJBALchTOF5hEoiDF42L5zq/CIOf8uRfrAeennuS6H18ATYiiHcmIIWuqOi4ZgcVy3ZPH81iczQ0A5UKMHUN+IXq0sCQQCoecNbkS2KTbWy+/Vgf+celRaM9CGGDfSNxVMIB/AaE730ZQ81YXdsoP1gSRElxPJclsb33AZzkuLCIp+GNb45AkAtQ/A201bHyBKcG00wYIrmgaSlwCtLJkrLj+LhvShnfWQTOQGfKwU2jEFFTWpmGaKuk2uNNoUlW3Rz9YI9n+07"
	key, err := ecrypto.RsaPublicKeyEncryptSymmetricKey([]byte("1122334455667788"), pk)
	fmt.Println(err)
	symmetricKey, err := ecrypto.RsaPrivateKeyDecryptSymmetricKey(key, sk)
	fmt.Println(err)
	fmt.Println(string(symmetricKey))
}
