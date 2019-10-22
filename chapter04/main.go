package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter04/router"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
