package main

import (
	"bytes"
	"fmt"
	"github.com/cattle2002/easycrypto/ecrypto"
	"io"
	"os"
)

func CreateFile(filename string, data []byte) error {
	create, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	reader := bytes.NewReader(data)
	_, err = io.Copy(create, reader)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func main() {
	file, err := os.ReadFile("D:\\workdir\\DIMSMonitorPlat\\productData\\aes小幸运歌曲数据产品.txt")
	//[203 73 240 191 63 27 234 142 35 68 172 102 223 145 186 209]
	key := []byte{203, 73, 240, 191, 63, 27, 234, 142, 35, 68, 172, 102, 223, 145, 186, 209}
	de, err := ecrypto.AesDecrypt(file, key, []byte("1234567812345678"))
	fmt.Println(err)
	CreateFile("ok.txt", de)
}
