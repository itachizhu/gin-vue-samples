package service

import (
	"github.com/itachizhu/gin-vue-samples/chapter10-02/logger"
	"github.com/itachizhu/gin-vue-samples/chapter10-02/model"
	"github.com/jinzhu/gorm"
)

func init()  {
	if mysql, err := gorm.Open("mysql", "root:12345678@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
		logger.Errorf("connection database error, %v", err)
	} else if nil != mysql {
		defer func(mysql *gorm.DB) {
			if err := mysql.Close(); nil != err {
				logger.Errorf("Disconnect from database failed: %v", err)
			}
		}(mysql)

		logger.Debug("connect success!")

		// 需要自动创建表的模型
		if err = mysql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Product{}).Error; nil != err {
			logger.Errorf("auto migrate tables failed: %v", err)
		}
	}
}
