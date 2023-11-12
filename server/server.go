package server

import (
	"DIMSMonitorPlat/config"
	"DIMSMonitorPlat/service"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

//var CoreServerConn *websocket.Conn
//var FrontEndConn *websocket.Conn

type Server struct {
	ip            string
	port          int
	coreServerUrl string
}

func NewServer(conf *config.Config) *Server {
	return &Server{ip: conf.Local.Host, port: conf.Local.Port, coreServerUrl: conf.PlatformUrl}
}

func (s *Server) Run() {
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
		<-sigChan
		os.Exit(2)
	}()
	addr := s.ip + ":" + strconv.Itoa(s.port)
	fmt.Println("程序运行地址：", addr)

	http.HandleFunc("/api/v1/calc", service.Calc)
	//http.HandleFunc("/api/v1/encrypt/symmetric/key", handle.EncryptSymmetricKey)
	//http.HandleFunc("/api/v1/decrypt/symmetric/key", handle.DecryptCipherSymmetricKey)
	http.HandleFunc("/api/v1/algo/list", service.Algo)
	http.HandleFunc("/api/v1/algo/register", service.Algo)
	http.HandleFunc("/api/v1/status", service.Status)
	http.HandleFunc("/api/v1/algo/keyword", service.KeyWord)
	http.HandleFunc("/api/v1/product/use", service.Use)
	http.HandleFunc("/te", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
		return
	})
	fmt.Println("server is running")
	err := http.ListenAndServe(":5517", nil)
	if err != nil {
		panic(err)
	}
}
