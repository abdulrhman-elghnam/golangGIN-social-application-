package app

import (
	"github.com/gin-gonic/gin"
)

func RegisterMainRoutes(rg *gin.RouterGroup) {
	mainGroup := rg.Group("/")
	{
		mainGroup.GET("/", MainController)
	}
}
