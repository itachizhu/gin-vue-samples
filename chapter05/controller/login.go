package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter05/model"
	"net/http"
)

func LoginGetAction(context *gin.Context) {
	context.HTML(http.StatusOK, "login.gohtml", nil)
}

func LoginPostAction(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		// 把错误信息返回到前端
		user.UserName = err.Error()
		context.HTML(http.StatusOK, "results.gohtml", user)
		return
	}
	context.HTML(http.StatusOK, "results.gohtml", user)
}

func LoginPutAction(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		return
	}
	context.JSON(http.StatusOK, user)
}
