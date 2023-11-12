package service

import (
	"DIMSMonitorPlat/handle"
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/model"
	"DIMSMonitorPlat/protocol"
	"DIMSMonitorPlat/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 前端使用数据产品
func serviceCalcError(errmsg string) []byte {
	response := protocol.HttpCalcResponse{
		Cmd:      string(protocol.Calc),
		RetCode:  protocol.ErrCode,
		ErrorMsg: errmsg,
		Payload:  protocol.HttpCalcResponsePayload{},
	}
	marshal, _ := json.Marshal(response)
	return marshal
}
func serviceCalcSuccess(response *protocol.HttpCalcResponse) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(response)
	if err != nil {
		return nil, err
	}
	marshal := buf.Bytes()
	return marshal, nil
}
func Before(calcReq protocol.HttpCalcRequest) error {
	//todo 向核心服务器上报,本地使用次数减1
	count, err := model.ReduceLocalUseNumber(calcReq.Payload.ProductID, calcReq.Payload.ProductName)
	if err != nil {
		log.Errorf("ReduceLocalUseNumber error:%s", err.Error())
		return err
	}
	err = handle.LocalUseSync(calcReq.Payload.ProductID, calcReq.Payload.Buyer, count)
	if err != nil {
		return err
	}
	return nil
}
func Calc(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		calcError := serviceCalcError(err.Error())
		w.Write(calcError)
		return
	}
	var calcReq protocol.HttpCalcRequest
	err = json.Unmarshal(body, &calcReq)
	if err != nil {
		calcError := serviceCalcError("非法协议")
		w.Write(calcError)
		return
	}
	fmt.Println("Request----------:", calcReq)
	//todo 先去sqlite查找是否有着一条数据 如果拥有这条数据 直接使用,如果没有这条数据去平台下载限制信息
	find, pd, err := model.Find(calcReq.Payload.ProductID, calcReq.Payload.ProductName)
	if err != nil {
		log.Errorf("查找数据产品ID:%d数据产品名字:%s的限制信息错误:%s", calcReq.Payload.ProductID, calcReq.Payload.ProductName, err.Error())
		calcError := serviceCalcError("查找数据产品限制信息失败")
		w.Write(calcError)
		return
	} else {
		if find == true {
			if pd.CanUseNumberLocalFlag == true {
				if pd.CanUseNumberLocal > 0 {
					//todo 处理数据产品
					calc, err := handle.Calc(&calcReq)
					if err != nil {
						calcError := serviceCalcError(err.Error())
						w.Write(calcError)
						return
					}
					success, err := serviceCalcSuccess(calc)
					if err != nil {
						calcError := serviceCalcError(err.Error())
						w.Write(calcError)
						return
					} else {
						w.Write(success)
						err := Before(calcReq)
						if err != nil {
							log.Errorf("Before ERROR:%s", err.Error())
						}
						return
					}
				} else {
					calcError := serviceCalcError("数据产品没有可使用次数")
					w.Write(calcError)
					return
				}
			} else {
				f := utils.GetCurrentMillSecond(pd.CanUseTimeLocal)
				if f == true {
					calcError := serviceCalcError("当前数据产品已过期")
					w.Write(calcError)
					return
				} else {
					calc, err := handle.Calc(&calcReq)
					if err != nil {
						calcError := serviceCalcError(err.Error())
						w.Write(calcError)
						return
					}
					success, err := serviceCalcSuccess(calc)
					if err != nil {
						calcError := serviceCalcError(err.Error())
						w.Write(calcError)
						return
					} else {
						w.Write(success)
						err := Before(calcReq)
						if err != nil {
							log.Errorf("Before Error:%s", err.Error())
						}
						return
					}
				}
			}
		} else {
			localUserNumber, localUseTime, err := handle.ProductLimitRequest(calcReq.Payload.ProductID, calcReq.Payload.Buyer, calcReq.Payload.ProductName)
			if err != nil {
				log.Errorf("Get %s limit information error:%s", err.Error())
				calcError := serviceCalcError(err.Error())
				w.Write(calcError)
				return
			}
			fmt.Println("返回结果：", localUserNumber, localUseTime)
			if utils.GetCurrentMillSecond(localUseTime) {
				fmt.Println("Seller", calcReq.Payload.Seller)
				err := model.Create(calcReq.Payload.ProductID, calcReq.Payload.ProductName, calcReq.Payload.Seller, true, false, localUserNumber, localUseTime)
				if err != nil {
					log.Errorf("更新产品限制信息到数据库失败:%s", err.Error())
					calcError := serviceCalcError("更新产品限制信息到数据库失败")
					w.Write(calcError)
					return
				}
			} else {
				err := model.Create(calcReq.Payload.ProductID, calcReq.Payload.ProductName, calcReq.Payload.Seller, false, true, localUserNumber, localUseTime)
				if err != nil {
					fmt.Println("为什么插不进去", err)
					log.Errorf("更新产品限制信息到数据库失败:%s", err.Error())
					calcError := serviceCalcError("更新产品限制信息到数据库失败")
					w.Write(calcError)
					return
				}
			}
			//
			response, err := handle.Calc(&calcReq)
			if err != nil {
				calcError := serviceCalcError(err.Error())
				w.Write(calcError)
				return
			}
			success, err := serviceCalcSuccess(response)
			if err != nil {
				calcError := serviceCalcError(err.Error())
				w.Write(calcError)
				return
			} else {
				w.Write(success)
				err := Before(calcReq)
				if err != nil {
					log.Errorf("Before Error:%s", err.Error())
				}
				return
			}
		}
	}
}
