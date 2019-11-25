package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter08-02/logger"
	"github.com/itachizhu/gin-vue-samples/chapter08-02/router"
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

	store := memstore.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysession", store))

	engine.Use(auditLogger())
	engine.Use(func(context *gin.Context) {
		context.Next()
	})
	engine.LoadHTMLGlob("templates/*")
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
