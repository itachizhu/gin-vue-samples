package model

type User struct {
	UserName string `json:"userName" form:"userName" binding:"required"` // 校验用户名非空
	Password string `json:"password" form:"password" binding:"required"` // 校验密码非空
}
