package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter03-2/router"
)

func main() {
	engine := gin.Default()
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
