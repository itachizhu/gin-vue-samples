package model

type Product struct {
	ID    uint64 `gorm:"primary_key" json:"id" form:"id"`                                      // 主键，自增长
	Code  string `gorm:"size:128" json:"code,omitempty" form:"productCode" binding:"required"` // 商品编码
	Name  string `gorm:"size:256" json:"name,omitempty" form:"productName" binding:"required"` // 商品名称
	Stock uint   `gorm:"type:bigint(20)" json:"stock" form:"productStock"`                     // 库存
}
