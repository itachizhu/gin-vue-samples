package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexJsonAction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"name": "itachi", "content": "Hello World!"})
}

func IndexXmlAction(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{"name": "itachi", "content": "Hello World!"})
}

func IndexYamlAction(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{"name": "itachi", "content": "Hello World!"})
}

func IndexImageAction(context *gin.Context) {
	context.File("./static/sorry.gif")
}

func IndexFileAction(context *gin.Context) {
	context.File("./static/demo.zip")
}
