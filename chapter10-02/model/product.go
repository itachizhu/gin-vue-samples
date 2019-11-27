package model

type Product struct {
	ID         uint64     `gorm:"primary_key" json:"id"` // 主键，自增长
	Code       string     `gorm:"size:128" json:"code,omitempty"` // 商品编码
	Name       string     `gorm:"size:256" json:"name,omitempty"` // 商品名称
	Stock      uint       `gorm:"type:bigint(20)" json:"stock"`   // 库存
}