package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter08-02/logger"
	"github.com/itachizhu/gin-vue-samples/chapter08-02/model"
	"net/http"
)

func LoginGetAction(context *gin.Context) {
	logger.Debug("您进入了登录页面")

	session := sessions.Default(context)
	if session.Get("user_name") == nil {
		context.HTML(http.StatusOK, "login.gohtml", nil)
	} else {
		userName := session.Get("user_name").(string)
		if userName == "" {
			context.HTML(http.StatusOK, "login.gohtml", nil)
		} else {
			var user model.User
			user.UserName = userName
			context.HTML(http.StatusOK, "results.gohtml", user)
		}
	}
}

func LoginPostAction(context *gin.Context) {
	logger.Info("你点击了登录按钮。")
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		logger.Errorf("登录错误，错误信息：%v", err)
		// 把错误信息返回到前端
		user.UserName = "用户名不能为空"
		user.Password = "密码不能为空"
		context.HTML(http.StatusOK, "results.gohtml", user)
		return
	}

	session := sessions.Default(context)
	session.Set("user_name", user.UserName)
	_ = session.Save()

	context.HTML(http.StatusOK, "results.gohtml", user)
}

func LoginPutAction(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		return
	}
	context.JSON(http.StatusOK, user)
}
