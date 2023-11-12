package main

import (
	"DIMSMonitorPlat/config"
	"DIMSMonitorPlat/handle"
	"DIMSMonitorPlat/log"
	"DIMSMonitorPlat/model"
	"DIMSMonitorPlat/ominio"
	"DIMSMonitorPlat/server"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

var CH chan int

func main() {
	CH = make(chan int)
	conf := &config.Config{}
	err := config.NewConfig("config.json", conf)
	if err != nil {
		log.Errorf("读取配置文件错误:%s", err.Error())
		os.Exit(1)
	}
	config.Conf = conf
	err = ominio.NewMinioClient()
	if err != nil {
		log.Infof("连接Minio错误:%s", err.Error())
	}
	username := conf.Local.User
	password := conf.Local.Password
	err = model.NewProductFile()
	err = model.OpenLimitDB()
	err = model.NewAlgoFile()
	err = model.OpenAlgoDB()
	// 日志初始化
	err = log.New(&log.Config{
		DebugMode:  conf.LogConfig.DebugMode,
		MaxSize:    conf.LogConfig.MaxSize,
		MaxAge:     conf.LogConfig.MaxAge,
		MaxBackups: conf.LogConfig.MaxBackups,
		Filename:   conf.LogConfig.Filename,
		Stdout:     true,
	})

	go func() {
		ctx, canel := context.WithTimeout(context.Background(), time.Second*10)
		defer canel()
		conn, _, err := websocket.DefaultDialer.DialContext(ctx, conf.PlatformUrl, nil)
		if err != nil {
			log.Errorf("连接核心服务器失败:%s", err.Error())
			CH <- 1
			return
		}
		handle.CoreServerConnStatus = true
		log.Info("连接核心服务器成功")
		handle.CoreServerConn = conn
		CH <- 1
		go handle.MonitorHandle(username, password)

	}()
	go func() {
		for {
			time.Sleep(time.Second * 20)
			fmt.Println("当前协程数量:", runtime.NumGoroutine())
		}
	}()
	go func() {
		select {
		case <-CH:
			for {
				//todo 判断 CoreServerStatus
				if handle.CoreServerConnStatus == false {
					ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
					conn, _, err := websocket.DefaultDialer.DialContext(ctx, conf.PlatformUrl, nil)
					if err != nil {
						log.Errorf("reconnect CoreServer error:%s", err.Error())
					} else {
						log.Info("reconnect CoreServer Success")
						handle.CoreServerConnStatus = true
						handle.CoreServerConn = conn
					}
				} else {
					time.Sleep(time.Second * 20)
				}
			}
		}
	}()
	//for i := 1; i < 15; i++ {
	//	time.Sleep(time.Second * 3)
	//	request, i, err := handle.ProductLimitRequest(86, "zcl01", "在线产品测试1106")
	//	fmt.Println(err)
	//	log.Info("限制信息：%d,%d", request, i)
	//
	//}
	newServer := server.NewServer(conf)
	newServer.Run()

}
