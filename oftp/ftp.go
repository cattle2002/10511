package oftp

import (
	"DIMSMonitorPlat/config"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"os"
	"time"
)

type Ftp struct {
	ftp     *ftp.ServerConn
	rootDir string
}

func Open() (*Ftp, error) {
	var conf config.Config
	err := config.NewConfig("./config.json", &conf)
	if err != nil {
		return nil, err
	}
	address := fmt.Sprintf("%s:%d", conf.FileServer.Host, conf.FileServer.Port)
	f, err := ftp.Dial(address, ftp.DialWithTimeout(time.Duration(10)*time.Second))
	if err != nil {
		return nil, err
	}
	err = f.Login(conf.FileServer.User, conf.FileServer.Password)
	if err != nil {
		return nil, err
	}
	return &Ftp{
		ftp:     f,
		rootDir: conf.FileServer.RootDir,
	}, nil
}

//func Upload() error {
//
//}

func Download(fileName string) ([]byte, error) {
	FtpConn, err := Open()
	if err != nil {
		return nil, err
	}
	err = FtpConn.ftp.ChangeDir("/product")
	if err != nil {
		return nil, err
	}
	size, _ := FtpConn.ftp.FileSize(fileName)
	fmt.Println("服务器文件大小:", size)
	// 打开远程文件
	r, err := FtpConn.ftp.Retr(fileName)
	if err != nil {
		fmt.Println("FTP retr error:", err)

	}
	defer r.Close()

	//获取本地存放product的路径
	dir := config.GetProductDir()
	// 创建本地文件
	localFile, err := os.Create(dir + fileName)
	if err != nil {
		fmt.Println("Local file creation error:", err)

	}
	defer localFile.Close()

	// 复制文件内容
	nw, err := io.Copy(localFile, r)
	fmt.Println("copy size:", nw)
	FtpConn.ftp.Quit()
	return nil, nil
}

func Upload(fileName string) error {
	FtpConn, err := Open()
	if err != nil {
		return err
	}
	defer FtpConn.ftp.Quit()
	err = FtpConn.ftp.ChangeDir("/product")
	if err != nil {
		return err
	}
	//打开本地文件
	open, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer open.Close()
	//创建远程文件
	err = FtpConn.ftp.Stor(fileName, open)
	return err
}
