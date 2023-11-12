package handle

import (
	"DIMSMonitorPlat/config"
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/ominio"
	"DIMSMonitorPlat/protocol"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/cattle2002/easycrypto/ecrypto"
	"strings"
	"time"
)

// Calc 与前端处理 解密数据产品 返回前端数据产品下载的url
func Calc(request *protocol.HttpCalcRequest) (*protocol.HttpCalcResponse, error) {
	//todo 先去sqlite查询是否拥有这条记录,如果没有向核心服务器获取这条数据产品的限制信息, 如果有的话查看本地使用次数和本地使用时间是否到期，
	if request.Payload.HaveData == false {
		if request.Payload.CertificateAlgorithmTypeAlgorithmType == config.Conf.KeyPair.Algorithm {
			//读取本地配置的私钥
			PrivatePem, err := config.GetPrivateKeyPem()
			if err != nil {
				log.Errorf("用户私钥未找到:%s", err.Error())
				return nil, errors.New(string(protocol.UserPrivateKeyNotFound))
			}
			//fmt.Println(PrivatePem)
			symmetricKey, err := ecrypto.RsaPrivateKeyDecryptSymmetricKey(request.Payload.CipherSymmetricKey, PrivatePem)
			if err != nil {
				log.Errorf("私钥解密加密对称密钥失败:%s", err.Error())
				return nil, errors.New(string(protocol.DecryptCipherSymmetricKey))
			}
			//替换成minio下载数据产品的api
			prodductUrl := strings.Replace(request.Payload.ProductUrl, "product/", "", -1)

			fb, err := ominio.DownloadBinary(context.Background(), config.Conf.Minio.ProductBucket, prodductUrl)
			if err != nil {
				log.Error("下载加密数据数据产品失败", err.Error())
				return nil, errors.New("没有找到数据产品")
			}

			if request.Payload.SymmetricKeyAlgorithmType == "aes" {
				//decrypt, err := ecrypto.AesDecrypt(fb, symmetricKey, []byte("1234567812345678"))
				decrypt, err := ecrypto.AesDecrypt(fb, symmetricKey, []byte("0000000000000000"))
				if err != nil {

					log.Errorf("Aes解密加密数据产品失败:%s", err.Error())
					return nil, errors.New("加密数据产品解密失败")
				}
				reader := bytes.NewReader(decrypt)
				err = ominio.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, prodductUrl, reader, int64(len(decrypt)))
				if err != nil {
					log.Errorf("上传解密后的数据产品失败:%s", err.Error())
					return nil, errors.New("上传解密后的数据产品失败")
				}
				url, err := ominio.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, prodductUrl, time.Second*60*30)
				if err != nil {
					log.Errorf("获取解密数据产品的下载链接失败:%s", err.Error())
					return nil, errors.New("获取解密数据产品的下载链接失败")
				}
				return &protocol.HttpCalcResponse{
					Cmd:      string(protocol.CalcRet),
					RetCode:  protocol.SuccessCode,
					ErrorMsg: "",
					Payload: protocol.HttpCalcResponsePayload{
						ID:       request.Payload.ID,
						HaveData: false,
						Data:     nil,
						Url:      url,
					},
				}, nil
			}
			if request.Payload.SymmetricKeyAlgorithmType == "sm4" {
				//todo 先base64 解码 错误bug 看DIMSCASOV3 Computeso sm4
				decodeString, err := base64.StdEncoding.DecodeString(string(fb))
				if err != nil {
					log.Errorf("Base64解码失败:%s", err.Error())
					return nil, errors.New("加密数据产品解密失败")
				}
				decrypt, err := ecrypto.Sm4Decrypt(decodeString, symmetricKey)
				if err != nil {
					log.Errorf("SM4 解密失败:%s", err.Error())
					return nil, errors.New("加密数据产品解密失败")
				}
				reader := bytes.NewReader(decrypt)
				err = ominio.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, prodductUrl, reader, int64(len(decrypt)))
				if err != nil {
					log.Errorf("上传解密数据产品失败:%s", err.Error())
					return nil, errors.New("上传解密数据产品失败")
				}
				url, err := ominio.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, prodductUrl, time.Second*60*10)
				if err != nil {
					log.Errorf("获取解密数据产品链接失败:%s", err.Error())
					return nil, errors.New("获取解密数据产品链接失败")
				}
				return &protocol.HttpCalcResponse{
					Cmd:      string(protocol.CalcRet),
					RetCode:  protocol.SuccessCode,
					ErrorMsg: "",
					Payload: protocol.HttpCalcResponsePayload{
						ID:       request.Payload.ID,
						HaveData: false,
						Data:     nil,
						Url:      url,
					},
				}, nil
			}
			return nil, errors.New("不支持请求算法")
		} else {
			return nil, errors.New("与当前证书算法不匹配")
		}
	}
	return nil, errors.New("平台还未支持数据产品携带在报文中")
}
