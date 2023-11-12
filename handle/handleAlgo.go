package handle

import (
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/model"
	"DIMSMonitorPlat/protocol"
)

func ConvertAlgModelToHttpModel(algoModel []model.AlgoModel) *[]protocol.HttpAlgoModel {
	var httpModel []protocol.HttpAlgoModel
	for i := range algoModel {
		var m protocol.HttpAlgoModel
		m.AlgoType = algoModel[i].AlgoType
		m.FileType = algoModel[i].FileType
		m.AlgoPosition = algoModel[i].AlgoPosition
		m.RunningEnv = algoModel[i].RunningEnv
		m.AlgoID = algoModel[i].AlgoID
		httpModel = append(httpModel, m)
	}
	return &httpModel
}
func Algo(req *protocol.HttpListAlgoRequest) (*protocol.HttpListAlgoResponse, error) {

	page, err := model.FindAlgoPage(req.PageNum, req.PageSize)
	if err != nil {
		log.Errorf("查找本地算法失败:%s", err.Error())
		return nil, err
	}
	response := protocol.HttpListAlgoResponse{
		Code: protocol.SuccessCode,
		Msg:  "",
		Data: ConvertAlgModelToHttpModel(page),
	}
	return &response, nil
}

//func AlgoListByKeywords(key string)
