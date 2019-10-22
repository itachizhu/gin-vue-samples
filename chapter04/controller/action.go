package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapGetAction(context *gin.Context) {
	// context.String(http.StatusOK, "hello gin get method")
	m := map[string]string{"title": "Map", "content": "Hello Map!"}
	context.HTML(http.StatusOK, "map.gohtml", m)
}

func VarGetAction(context *gin.Context) {
	// context.String(http.StatusOK, "hello gin get method")
	m := map[string]string{"title": "定义变量", "content": "Hello 变量!"}
	context.HTML(http.StatusOK, "var.gohtml", m)
}
