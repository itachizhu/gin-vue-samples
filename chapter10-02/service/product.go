package service

import (
	"github.com/itachizhu/gin-vue-samples/chapter10-02/logger"
	"github.com/itachizhu/gin-vue-samples/chapter10-02/model"
)

func SelectProduct() []model.Product {
	var products []model.Product
	if db == nil {
		logger.Debug("数据库连接对象为空，连接有问题")
		return products
	}
	if err := db.Find(&products).Error; err != nil {
		logger.Errorf("商品列表查询错误：%v", err)
	}
	return products
}

func GetProduct(id uint64) model.Product {
	var product model.Product
	if db == nil {
		logger.Debug("数据库连接对象为空，连接有问题")
		return product
	}
	if err := db.Where("`id`=?", id).First(&product).Error; err != nil {
		logger.Errorf("商品列表查询错误：%v", err)
	}
	return product
}

func SaveProduct(product model.Product) {
	if db == nil {
		logger.Debug("数据库连接对象为空，连接有问题")
		return
	}
	if err := db.Save(&product).Error; err != nil {
		logger.Error("保存商品失败： %v", err)
	}
}

func DeleteProduct(id uint64) {
	if db == nil {
		logger.Debug("数据库连接对象为空，连接有问题")
		return
	}

	if err := db.Where("`id`=?", id).Delete(model.Product{}).Error; err != nil {
		logger.Error("删除商品失败： %v", err)
	}
}
