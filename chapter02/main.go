package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
