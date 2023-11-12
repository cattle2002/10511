package handle

import "DIMSMonitorPlat/protocol"

func Status(req *protocol.HttpStatusReq) (*protocol.HttpStatusRes, error) {
	var res protocol.HttpStatusRes
	res.Code = protocol.SuccessCode
	res.Msg = "Running"
	return &res, nil
}
