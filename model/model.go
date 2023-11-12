package model

import (
	"DIMSMonitorPlat/config"
	"fmt"
	"gorm.io/gorm"
	"os"
	"runtime"
)

var DBLimit *gorm.DB
var DBAlgo *gorm.DB
var err error

// 判断当前项目的根目录下是否存在product_limit.db 文件
func NewProductFile() error {
	cdir := config.Conf.Local.CurrentDir
	fileName := "productdata_limit.db"
	if runtime.GOOS == "windows" {
		fmt.Println(cdir + "\\" + fileName)
		_, err := os.Stat(cdir + "\\" + fileName)
		if err == nil {
			return nil
		} else {
			_, err := os.Create(cdir + "\\" + fileName)
			if err != nil {
				return err
			} else {
				fmt.Println("初始化数据产品配置文件成功")
				return nil
			}
		}
	} else {
		_, err := os.Stat(cdir + "/" + fileName)
		if err != nil {
			return nil
		} else {
			_, err := os.Create(cdir + "/" + fileName)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}
func NewAlgoFile() error {
	cdir := config.Conf.Local.CurrentDir
	fileName := "Algo.db"
	if runtime.GOOS == "windows" {
		fmt.Println(cdir + "\\" + fileName)
		_, err := os.Stat(cdir + "\\" + fileName)
		if err == nil {
			return nil
		} else {
			_, err := os.Create(cdir + "\\" + fileName)
			if err != nil {
				return err
			} else {
				fmt.Println("初始化Algo文件成功")
				return nil
			}
		}
	} else {
		_, err := os.Stat(cdir + "/" + fileName)
		if err != nil {
			return nil
		} else {
			_, err := os.Create(cdir + "/" + fileName)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}
