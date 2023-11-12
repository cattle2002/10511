package handle

//
import (
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/packetmsg"
	"DIMSMonitorPlat/protocol"
	"DIMSMonitorPlat/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

var CoreServerConn *websocket.Conn
var CoreServerConnStatus bool

func MonitorHandle(username string, password string) {
	req := protocol.LoginReq{
		Cmd:     string(protocol.Login),
		Program: string(protocol.Monitor),
		Payload: protocol.LoginReqPayload{
			ID:        utils.MsgID(),
			TimeStamp: time.Now().Unix(),
			User:      username,
			Password:  password,
		},
	}
	marshal, _ := json.Marshal(req)
	packet, err := packetmsg.Packet(marshal)
	if err != nil {
		log.Errorf("packet error:%s", err.Error())

		return
	}
	for {
		time.Sleep(time.Second * 60)

		if CoreServerConnStatus == true {
			err := CoreServerConn.WriteMessage(websocket.TextMessage, packet)
			if err != nil {
				log.Errorf("CoreServer is offline:%s", err.Error())
				return
			}
			log.Info("CoreServer Heartbeat")

		} else {
			log.Error("CoreServer Heartbeat error:CoreServer is Offline")
			return
		}
	}
}

// func CoreServerHandle(username string, password string) {
//
//		for {
//			_, p, err := CoreServerConn.ReadMessage()
//			packet, err := packetmsg.Packet(marshal)
//			if err != nil {
//				log.Errorf("Packet msg  error:%s", err.Error())
//			}
//			err = CoreServerConn.WriteMessage(websocket.TextMessage, packet)
//			if err != nil {
//				log.Errorf("write msg  error:%s", err.Error())
//			}
//			_, p, err := CoreServerConn.ReadMessage()
//			if err != nil {
//				CoreServerConnStatus = false
//				return
//			} else {
//				uPacket, _ := packetmsg.UPacket(p)
//				var res protocol.LoginRes
//				err = json.Unmarshal(uPacket, &res)
//				if err != nil {
//					log.Errorf("marshal msg error:%s", err.Error())
//				}
//				fmt.Println("------------")
//				fmt.Println(res)
//				if res.RetCode == 0 {
//					fmt.Println(res)
//					log.Info("heartBeat ok")
//				}
//			}
//		}
//	}
func SymmetricKeyRequest(ProductID int64, Seller string, Buyer string) (string, string, error) {
	req := protocol.SymmetricReq{
		Cmd:     string(protocol.Symmetric),
		Program: string(protocol.Monitor),
		Payload: protocol.SymmetricReqPayload{
			ID:        utils.MsgID(),
			ProductID: ProductID,
			Seller:    Seller,
			Buyer:     Buyer,
		},
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		//log.Errorf("marshal error:%s", err.Error())
		return "", "", err
	}
	packet, err := packetmsg.Packet(marshal)
	err = CoreServerConn.WriteMessage(websocket.TextMessage, packet)
	if err != nil {
		//log.Errorf("write msg error:%s", err.Error())
		return "", "", err
	}
	_, p, err := CoreServerConn.ReadMessage()
	uPacket, err := packetmsg.UPacket(p)
	if err != nil {
		//log.Errorf("upacket error:%s",err.Error())
		return "", "", err
	}
	var res protocol.SymmetricRes
	err = json.Unmarshal(uPacket, &res)
	if err != nil {
		//log.Errorf("unmarshal error:%s",err.Error())
		return "", "", err
	}
	if res.RetCode != 0 {
		//log.Errorf("retcode error:%s",res.ErrorMsg)
		return "", "", err
	}
	return res.Payload.AlgoType, res.Payload.Key, nil
}
func ProductLimitRequest(productID int64, buyer string, productName string) (int64, int64, error) {
	req := protocol.LimitReq{
		Cmd:     string(protocol.Limit),
		Program: string(protocol.Monitor),
		Payload: protocol.LimitReqPayload{
			ID:          utils.MsgID(),
			Buyer:       buyer,
			ProductName: productName,
			ProductID:   productID,
		},
	}
	marshal, _ := json.Marshal(req)
	packet, err := packetmsg.Packet(marshal)
	if err != nil {
		log.Errorf("packet msg error:%s", err.Error())
		fmt.Println(err)
		return 0, 0, err
	}
	err = CoreServerConn.WriteMessage(websocket.TextMessage, packet)
	if err != nil {
		log.Errorf("write msg error:%s", err.Error())
		fmt.Println(err)
		return 0, 0, err
	}
	_, p, err := CoreServerConn.ReadMessage()
	if err != nil {
		log.Errorf("read msg error:%s", err.Error())
		fmt.Println("11")
		return 0, 0, err
	}
	uPacket, err := packetmsg.UPacket(p)
	if err != nil {
		log.Errorf("upacket msg error:%s", err.Error())
		return 0, 0, err
	}
	var res protocol.LimitRes
	err = json.Unmarshal(uPacket, &res)
	if err != nil {
		log.Errorf("unmarshal error:%s", err.Error())
		return 0, 0, err
	}
	fmt.Println("------------------", res)
	if res.RetCode == 0 {
		return res.Payload.Detail.CanUseNumberLocal, res.Payload.Detail.CanUseTimeLocal, nil
	} else {
		return 0, 0, errors.New(res.ErrorMsg)
	}
}

// 使用完数据产品的向平台上报
func LocalUseSync(ProductID int64, UserName string, CanUseNumberLocal int64) error {
	if CoreServerConnStatus == false {
		log.Error("核心服务器失去连接")
		return errors.New("核心服务器失去连接")
	}
	req := protocol.LocalCanUseSyncReq{
		Cmd:     string(protocol.LocalCanUse),
		Program: string(protocol.Monitor),
		Payload: protocol.LocalCanUseSyncReqPayload{
			ID:                utils.MsgID(),
			ProductID:         ProductID,
			UserName:          UserName,
			CanUseNumberLocal: CanUseNumberLocal,
		},
	}
	marshal, _ := json.Marshal(req)
	packet, err := packetmsg.Packet(marshal)
	if err != nil {
		log.Errorf("对消息封包出错:%s", err.Error())
		return errors.New("对消息封包出错")
	}
	err = CoreServerConn.WriteMessage(websocket.TextMessage, packet)
	if err != nil {
		log.Errorf("向核心服务发送消息错误:%s", err.Error())
		return errors.New("向核心服务发送消息错误")
	}
	_, p, err := CoreServerConn.ReadMessage()
	if err != nil {
		log.Errorf("接收核心服务器消息错误:%s", err.Error())
		return errors.New("接收核心服务器消息错误")
	}
	uPacket, err := packetmsg.UPacket(p)
	if err != nil {
		log.Errorf("对消息拆包失败:%s", err.Error())
		return errors.New("对消息拆包失败")
	}
	var res protocol.LocalCanUseSyncRes
	err = json.Unmarshal(uPacket, &res)
	if err != nil {
		log.Errorf("对消息反序列化失败:%s", err.Error())
		return errors.New("对消息反序列化失败")
	}
	if res.RetCode != 0 {
		log.Error(res.ErrorMsg)
		return errors.New(res.ErrorMsg)
	} else {
		return nil
	}
}
