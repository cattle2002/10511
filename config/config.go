package config

import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// var Conf.Local.CurrentDir string
// var Conf.Local.CurrentDir string
// var PrivateKeyFileName string
// var CertificateAlgoType string
var PublicKeyFileName string
var User string
var Conf *Config

func GetProductDir() string {
	if runtime.GOOS == "windows" {
		return Conf.Local.CurrentDir + "\\" + "productData\\"
	} else {
		return Conf.Local.CurrentDir + "/productData/"
	}
}
func GetPublicKeyPem() (string, error) {
	if runtime.GOOS == "windows" {
		publicFilePosition := Conf.Local.CurrentDir + "\\" + PublicKeyFileName
		file, err := os.ReadFile(publicFilePosition)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	} else {
		publicFilePosition := Conf.Local.CurrentDir + "/" + PublicKeyFileName
		file, err := os.ReadFile(publicFilePosition)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	}
}
func GetPrivateKeyPem() (string, error) {
	if runtime.GOOS == "windows" {
		privateFilePosition := Conf.Local.CurrentDir + "\\" + Conf.KeyPair.PrivateKeyPath
		file, err := os.ReadFile(privateFilePosition)
		if err != nil {
			return "", nil
		}
		return string(file), nil
	} else {
		fmt.Println("看看我2:")
		privateFilePosition := Conf.Local.CurrentDir + "/" + Conf.KeyPair.PrivateKeyPath
		file, err := os.ReadFile(privateFilePosition)
		if err != nil {
			return "", nil
		}
		fmt.Println("看看我2:")
		return string(file), nil
	}
}

type KeyPair struct {
	AutoConfig     bool   `json:"AutoConfig"`
	Algorithm      string `json:"Algorithm"`
	Bits           int    `json:"Bits"`
	PublicKeyPath  string `json:"PublicKeyPath"`
	PrivateKeyPath string `json:"PrivateKeyPath"`
}

type LoggerConfig struct {
	Filename   string `json:"Filename"`
	MaxSize    int    `json:"MaxSize"`
	MaxAge     int    `json:"MaxAge"`
	MaxBackups int    `json:"MaxBackups"`
	DebugMode  bool   `json:"DebugMode"`
	Stdout     bool   `json:"Stdout"`
}
type Local struct {
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	CurrentDir string `json:"CurrentDir"`
}

type FileServer struct {
	User     string `json:"User"`
	Password string `json:"Password"`
	Port     int    `json:"Port"`
	Host     string `json:"Host"`
	RootDir  string `json:"RootDir"`
}

type Minio struct {
	EndPoint        string `json:"EndPoint"`
	AccessKeyID     string `json:"AccessKeyID"`
	SecretAccessKey string `json:"SecretAccessKey"`
	UseSSL          bool   `json:"UseSSL"`
	ProductBucket   string `json:"ProductBucket"`
	ProductUpload   string `json:"ProductUpload"`
}
type Config struct {
	PlatformUrl string       `json:"PlatformUrl"`
	KeyPair     KeyPair      `json:"KeyPair"`
	LogConfig   LoggerConfig `json:"Logger"`
	Local       Local        `json:"Local"`
	FileServer  FileServer   `json:"FileServer"`
	Minio       Minio        `json:"Minio"`
}

func NewConfig(configPath string, v interface{}) (err error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, v); err != nil {
		return err
	}

	return nil
}
