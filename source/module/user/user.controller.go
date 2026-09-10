package user


import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	UserRoute := rg.Group("/user")
	{
		UserRoute.GET("/", GetUser)
	}
}