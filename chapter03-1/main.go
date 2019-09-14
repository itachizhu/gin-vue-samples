package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 添加 Get 请求路由
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin get method")
	})
	// 添加 Post 请求路由
	engine.POST("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})
	// 添加 Put 请求路由
	engine.PUT("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin put method")
	})
	// 添加 Delete 请求路由
	engine.DELETE("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin delete method")
	})
	// 添加 Patch 请求路由
	engine.PATCH("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin patch method")
	})
	// 添加 Head 请求路由
	engine.HEAD("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin head method")
	})
	// 添加 Options 请求路由
	engine.OPTIONS("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin options method")
	})
	// 添加处理任意方法的请求路由
	engine.Any("/hi", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin any method")
	})
	// 使用Handle方法添加 Get 请求路由
	engine.Handle("GET", "/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin handle get method")
	})

	// 获取url参数
	engine.GET("/user", func(context *gin.Context) {
		name := context.Query("name")
		context.String(http.StatusOK, "hello "+name)
	})
	// 获取表单参数
	engine.POST("/user", func(context *gin.Context) {
		name := context.PostForm("name")
		context.String(http.StatusOK, "hello "+name)
	})
	// 获取请求body参数
	engine.PUT("/user", func(context *gin.Context) {
		var m map[string]string
		if err := context.BindJSON(&m); err != nil {
			context.String(http.StatusInternalServerError, "error data!")
			return
		}
		context.String(http.StatusOK, "hello "+m["name"])
	})
	// 获取路径参数
	engine.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello "+name)
	})

	// 路由分组
	admin := engine.Group("/admin")
	{
		admin.Any("/hello", func(context *gin.Context) {
			context.String(http.StatusOK, "hello we are admin group!")
		})
	}

	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
