package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter08-01/logger"
	"github.com/itachizhu/gin-vue-samples/chapter08-01/router"
	"time"
)

func auditLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path

		context.Next()

		logger.Infof("%s | %3d | %15s | %-7s | %s\n%s", start.Format("2006/01/02 - 15:04:05"),
			context.Writer.Status(),
			context.ClientIP(),
			context.Request.Method,
			path,
			context.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}

func main() {
	engine := gin.Default()
	engine.Use(auditLogger())
	engine.Use(func(context *gin.Context) {
		context.Next()
	})
	engine.LoadHTMLGlob("templates/*")
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
