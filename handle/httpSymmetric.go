package handle

import (
	"DIMSMonitorPlat/config"
	"DIMSMonitorPlat/protocol"
	"encoding/json"
	"fmt"
	"github.com/cattle2002/easycrypto/ecrypto"
	"io"
	"net/http"
)

func HttpResponseCryptoError(errmsg string) []byte {
	response := &protocol.HttpCryptoRes{
		Cmd:      string(protocol.MonitorRet),
		RetCode:  protocol.ErrCode,
		ErrorMsg: errmsg,
		Payload:  protocol.HttpCryptoResPayload{},
	}
	marshal, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	return marshal
}
func EncryptSymmetricKey(w http.ResponseWriter, r *http.Request) {
	all, err := io.ReadAll(r.Body)
	if err != nil {
		//todo
		fmt.Println(err)
	}
	var req protocol.HttpCryptoReq
	err = json.Unmarshal(all, &req)
	if err != nil {
		fmt.Println(err)
		responseError := HttpResponseCryptoError("消息格式不正确")
		w.Write(responseError)
		return
	}
	if len(req.Payload.SymmetricKey) != 16 {
		responseError := HttpResponseCryptoError("密钥长度不对")
		w.Write(responseError)
		return
	}
	if req.Payload.User == config.User {
		//todo 读取用户的公钥
		pem, err := config.GetPublicKeyPem()
		if err != nil {
			//todo 返回管理服务器相应
			fmt.Println(err)
			responseError := HttpResponseCryptoError("内部错误")
			w.Write(responseError)
			return
		}
		if config.Conf.KeyPair.Algorithm == "rsa" {
			key, err := ecrypto.RsaPublicKeyEncryptSymmetricKey([]byte(req.Payload.SymmetricKey), pem)
			if err == nil {
				res := protocol.HttpCryptoRes{
					Cmd:      string(protocol.CryptoRet),
					RetCode:  protocol.SuccessCode,
					ErrorMsg: "",
					Payload:  protocol.HttpCryptoResPayload{CipherSymmetricKey: key, AlgoType: config.Conf.KeyPair.Algorithm},
				}
				marshal, _ := json.Marshal(res)
				w.Write(marshal)
				return
			} else {
				responseError := HttpResponseCryptoError("内部错误")
				w.Write(responseError)
				return
			}
		}
		if config.Conf.KeyPair.Algorithm == "sm2" {
			key, err := ecrypto.Sm2PublicKeyEncryptSymmetricKey([]byte(req.Payload.SymmetricKey), pem)
			if err != nil {
				marshal := HttpResponseCryptoError("内部错误")
				w.Write(marshal)
				return
			}
			res := protocol.HttpCryptoRes{
				Cmd:      string(protocol.CryptoRet),
				RetCode:  protocol.SuccessCode,
				ErrorMsg: "",
				Payload:  protocol.HttpCryptoResPayload{CipherSymmetricKey: key, AlgoType: config.Conf.KeyPair.Algorithm},
			}
			marshal, _ := json.Marshal(res)
			w.Write(marshal)
			return
		}
	} else {
		responseError := HttpResponseCryptoError("用户名不正确")
		w.Write(responseError)
		return
	}
}

//type Jiemi struct {
//	Key string `json:"Key"`
//}

func DecryptCipherSymmetricKey(w http.ResponseWriter, r *http.Request) {
	all, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var req protocol.HttpDecryptReq
	//var jiemi Jiemi
	err = json.Unmarshal(all, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	pem, err := config.GetPrivateKeyPem()
	if err != nil {
		fmt.Println(err)
		return
	}
	var res protocol.HttpDecryptRes
	key, err := ecrypto.RsaPrivateKeyDecryptSymmetricKey(req.Payload.CipherSymmetricKey, pem)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Payload.SymmetricKey = string(key)
	marshal, _ := json.Marshal(res)
	w.Write(marshal)
}
