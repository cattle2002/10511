package service

import (
	"DIMSMonitorPlat/handle"
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/model"
	"DIMSMonitorPlat/protocol"
	"encoding/json"
	"io"
	"net/http"
)

func serviceAlgoError(msg string) []byte {
	rs := protocol.HttpListAlgoResponse{
		Code: protocol.ErrCode,
		Msg:  msg,
		Data: nil,
	}
	marshal, _ := json.Marshal(rs)
	return marshal
}
func serviceAlgoSuccess(algos []protocol.HttpAlgoModel) []byte {
	response := protocol.HttpListAlgoResponse{
		Code: protocol.SuccessCode,
		Msg:  "",
		Data: &algos,
	}
	marshal, _ := json.Marshal(response)
	return marshal
}
func Algo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	body, err := io.ReadAll(r.Body)
	if err != nil {
		calcError := serviceCalcError(err.Error())
		w.Write(calcError)
		return
	}
	var req protocol.HttpListAlgoRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		algoError := serviceAlgoError("非法协议")
		w.Write(algoError)
		return
	}
	algos, err := handle.Algo(&req)
	if err != nil {
		calcError := serviceCalcError(err.Error())
		w.Write(calcError)
		return
	}
	marshal, err := json.Marshal(algos)
	if err != nil {
		log.Errorf("marshal error:%s", err.Error())
		return
	}
	w.Write(marshal)
	return
}

type KeyWordsReq struct {
	Keyword string `json:"Keyword"`
}

func KeyWord(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	all, err := io.ReadAll(r.Body)
	if err != nil {
		algoError := serviceAlgoError("非法协议")
		w.Write(algoError)
		return
	}
	var kw KeyWordsReq
	err = json.Unmarshal(all, &kw)
	if err != nil {
		algoError := serviceAlgoError("非法协议")
		w.Write(algoError)
		return
	}
	words, err := model.FindAlgoByKeyWords(kw.Keyword)
	if err != nil {
		algoError := serviceAlgoError(err.Error())
		w.Write(algoError)
		return
	}
	httpModel := handle.ConvertAlgModelToHttpModel(words)
	success := serviceAlgoSuccess(*httpModel)
	w.Write(success)
	return
}
func Use(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

}
