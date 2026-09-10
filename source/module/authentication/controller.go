package authentication

import (
	"golang/source/common/middleware/pipe"
	"golang/source/common/structure"
	"golang/source/module/authentication/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(rg *gin.RouterGroup) {
	authenticationRoute := rg.Group("/authentication")

	{
		authenticationRoute.POST("/signup",
			pipe.Validate(func() any {
				return &dto.Signup{}
			}),
			func(ctx *gin.Context) {
				data := ctx.MustGet("body").(*dto.Signup)
				serviceFeedback, err := Signup(*data)
				if err != nil {
					structure.Fail(ctx , http.StatusInternalServerError , "ERROR" ,"error in authentication handler")
					return
				}
				structure.OK(ctx, 200, serviceFeedback)
			},
		)
	}
}