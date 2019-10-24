package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
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

func ConditionGetAction(context *gin.Context) {
	// context.String(http.StatusOK, "hello gin get method")
	m := map[string]int{"sum": rand.Intn(20)}
	context.HTML(http.StatusOK, "condition.gohtml", m)
}

func RangeGetAction(context *gin.Context) {
	m := map[string]string{"a": "A", "b": "B", "c": "C", "d": "D"}
	context.HTML(http.StatusOK, "range.gohtml", m)
}

func TemplateGetAction(context *gin.Context) {
	context.HTML(http.StatusOK, "main.gohtml", nil)
}
