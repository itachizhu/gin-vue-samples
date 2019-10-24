package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter04/router"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./static")
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
