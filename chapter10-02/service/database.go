package service

import (
	"github.com/itachizhu/gin-vue-samples/chapter10-02/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var db *gorm.DB

var once sync.Once

func Init() {
	once.Do(func() {
		var err error
		// <<username>>:<<password>>@tcp(<<ip>>:<<port>>)/<<dbname>>?charset=utf8mb4&parseTime=True&loc=Local
		if db, err = gorm.Open("mysql", "root:12345678@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
			logger.Errorf("connection database error, %v", err)
		} else {
			logger.Debug("connect success!")
			db.DB().SetMaxIdleConns(10)
			db.DB().SetMaxOpenConns(50)
			db.LogMode(true)
		}
	})
}

func Close() {
	if db != nil {
		if err := db.Close(); nil != err {
			logger.Errorf("Disconnect from database failed: %v", err)
		}
	}
	db = nil
}