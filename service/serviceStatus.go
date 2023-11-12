package service

import (
	"DIMSMonitorPlat/handle"
	"DIMSMonitorPlat/protocol"
	"encoding/json"
	"io"
	"net/http"
)

func serviceStatusError(errmsg string) []byte {
	var res protocol.HttpStatusRes
	res.Code = protocol.ErrCode
	res.Msg = errmsg
	marshal, _ := json.Marshal(res)
	return marshal
}
func serviceStatusSuccess(res *protocol.HttpStatusRes) []byte {
	marshal, _ := json.Marshal(res)
	return marshal
}

func Status(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set(" Access-Control-Request-Method", "*")
	//w.Header().Set("Access-Control-Request-Headers", "*")
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//w.Header().Set("content-type", "application/json")             //返回数据格式是json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	all, err := io.ReadAll(r.Body)
	if err != nil {
		statusError := serviceStatusError("非法协议")
		w.Write(statusError)
		return
	}
	var req protocol.HttpStatusReq
	err = json.Unmarshal(all, &req)
	if err != nil {
		statusError := serviceStatusError("非法协议")
		w.Write(statusError)
		return
	}
	status, err := handle.Status(&req)
	if err != nil {
		statusError := serviceStatusError("内部错误")
		w.Write(statusError)
		return
	}
	success := serviceStatusSuccess(status)
	w.Write(success)
	return
}
