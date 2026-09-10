package authentication


import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(rg *gin.RouterGroup) {
	AuthenticationRoute := rg.Group("/authentication")
	{
		AuthenticationRoute.POST("/", )
	}
}