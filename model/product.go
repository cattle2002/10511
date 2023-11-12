package model

import (
	"DIMSMonitorPlat/log"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Way:0 随便用 1仅限自己随便用 2监管使用 0.无限制 1.使用时长 2.访问下载时长 4.访问下载次数
// 8.使用次数 方式可以组合 如5表示有时长也有使用次数限制
type ProductDataLimitModel struct {
	ProductID             int64  `gorm:"column:product_id;unique;not null;primaryKey"`
	ProductName           string `gorm:"column:product_name;unique;not null"`
	Seller                string `gorm:"column:seller;not null"`
	CanUseNumberLocalFlag bool   `gorm:"column:number_flag;not null"`
	CanUseTimeLocalFlag   bool   `gorm:"column:time_flag;not null"`
	CanUseNumberLocal     int64  `gorm:"column:canuse_number;not null"`
	CanUseTimeLocal       int64  `gorm:"column:canuse_time;not null"`
}

func (this *ProductDataLimitModel) TableName() string {
	return "product_limit"
}

// 打开存放限制信息的sqlite
func OpenLimitDB() error {
	DBLimit, err = gorm.Open(sqlite.Open("productdata_limit.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DBLimit.AutoMigrate(&ProductDataLimitModel{})

	if err != nil {
		return err
	}
	return nil
}

// 限制信息
func Create(productID int64, productName string, seller string, numberFlag bool, timeFlag bool, canUseNumberLocal int64, canUseTimeLocal int64) error {
	err = OpenLimitDB()
	if err != nil {
		return err
	}
	// var m ProductDataLimitModel
	// m.Way = way
	tx := DBLimit.Create(&ProductDataLimitModel{ProductID: productID, ProductName: productName, Seller: seller, CanUseNumberLocalFlag: numberFlag, CanUseTimeLocalFlag: timeFlag, CanUseNumberLocal: canUseNumberLocal, CanUseTimeLocal: canUseTimeLocal})
	return tx.Error
}

// 限制信息
func Find(productID int64, productName string) (bool, *ProductDataLimitModel, error) {
	err = OpenLimitDB()
	if err != nil {
		log.Errorf("connect sqlite error:%s", err.Error())
		return false, nil, err
	}

	var pd ProductDataLimitModel
	tx := DBLimit.Debug().Where("product_id = ? AND product_name = ?", productID, productName).First(&pd)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, nil, nil
	}
	return true, &pd, nil
}

// 减少本地使用次数
func ReduceLocalUseNumber(productID int64, productName string) (int64, error) {
	DBLimit, err = gorm.Open(sqlite.Open("productdata_limit.db"), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	tx := DBLimit.Model(&ProductDataLimitModel{}).Where("product_id = ? and product_name = ?", productID, productName).Updates(map[string]interface{}{"CanUseLocalNumber": gorm.Expr("CanUseLocalNumber - ?", 1)})
	if tx.Error != nil {
		return 0, tx.Error
	}
	var m ProductDataLimitModel
	first := DBLimit.First(&m, "product_id=?", productID)
	return m.CanUseNumberLocal, first.Error
}
