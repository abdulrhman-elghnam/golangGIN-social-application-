package authentication

import (
	"golang/source/common/middleware/pipe"
	"golang/source/common/structure"
	"golang/source/module/authentication/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(rg *gin.RouterGroup) {
	AuthenticationRoute := rg.Group("/authentication")

		AuthenticationRoute.POST(
			"/signup",
			pipe.Validate(func() any {
				return &dto.Signup{}
			}),
			func(ctx *gin.Context) {
				result := SignupHandler(ctx)
				structure.OK(ctx, http.StatusCreated, result)
			},
		)	

		AuthenticationRoute.POST(
			"/login",
			pipe.Validate(func() any {
				return &dto.Signup{}
			}),
			func(ctx *gin.Context) {
				result := LoginHandler(ctx)
				structure.OK(ctx, http.StatusOK, result)
			},
		)
}
