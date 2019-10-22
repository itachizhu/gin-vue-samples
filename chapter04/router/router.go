package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itachizhu/gin-vue-samples/chapter04/controller"
)

func Route(engine *gin.Engine) {
	engine.GET("/", controller.IndexGetAction)
	engine.POST("/", controller.IndexPostAction)
	engine.PUT("/", controller.IndexPutAction)
	engine.DELETE("/", controller.IndexDeleteAction)

	engine.GET("/map", controller.MapGetAction)
	engine.GET("/var", controller.VarGetAction)
}
