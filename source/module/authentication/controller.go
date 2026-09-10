package authentication

import (
	"golang/source/common/middleware/pipe"
	"golang/source/common/structure"
	"golang/source/module/authentication/dto"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(rg *gin.RouterGroup) {
	authenticationRoute := rg.Group("/authentication")

	{
		authenticationRoute.POST(
			"/signup",
			pipe.Validate(func() any {
				return &dto.Signup{}
			}),
			func(ctx *gin.Context) {
				result := SignupHandler(ctx)
				structure.OK(ctx, 201, result)
			},
		)
	}
}
