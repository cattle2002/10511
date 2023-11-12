package model

import (
	"DIMSMonitorPlat/log"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AlgoModel struct {
	gorm.Model
	AlgoID       int64  `gorm:"column:algoID;not null;unique"`
	AlgoType     string `gorm:"column:algoType;not  null"`
	AlgoPosition string `gorm:"column:algoPosition;not null"`
	AlgoName     string `gorm:"column:algoName;not null"`
	RunningEnv   string `gorm:"column:runningEnv;not  null"`
	FileType     string `gorm:"column:fileType;not null"`
	AlgoMore     string `gorm:"column:algoMore;not null"`
	Details      string `gorm:"column:details;not null"`
	Algos        string `gorm:"column:algos;not null"`
	ProcessOrder string `gorm:"column:processOrder;not null"`
}
type Algo struct {
	ProcessID        int    `json:"ProcessID"`
	AlgoType         string `json:"AlgoType"`
	AlgoPosition     string `json:"AlgoPosition"`
	DependencyBefore string `json:"DependencyBefore"`
	NeedArgs         string `json:"NeedArgs"`
	FuncName         string `json:"FuncName"`
	FuncArgs         string `json:"FuncArgs"`
	Output           bool   `json:"Output"`
	OutputPosition   string `json:"OutputPosition"`
	Value            string `json:"Value"`
}

func (this *AlgoModel) TableName() string {
	return "algos"
}
func OpenAlgoDB() error {
	DBAlgo, err = gorm.Open(sqlite.Open("Algo.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBAlgo.AutoMigrate(&AlgoModel{})
	if err != nil {
		return err
	}
	return nil
}

// 查询
func FindAlgoPage(pageNum int, pageSize int) ([]AlgoModel, error) {
	var AlgoModels []AlgoModel
	err := OpenAlgoDB()
	if err != nil {
		fmt.Println(err)
	}
	tx := DBAlgo.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&AlgoModels)

	return AlgoModels, tx.Error
}
func FindAlgoByKeyWords(keyWord string) ([]AlgoModel, error) {
	err := OpenAlgoDB()
	if err != nil {
		log.Errorf("连接Sqlite数据库失败:%s", err.Error())
		return nil, err
	}
	fmt.Println("----------k,", keyWord)
	var algos []AlgoModel
	tx := DBAlgo.Debug().Where("algoName LIKE ?", "%"+keyWord+"%").Find(&algos)
	if tx.Error != nil {
		log.Errorf("查找算法信息失败:%s", err.Error())
		return nil, err
	}
	return algos, nil
}
