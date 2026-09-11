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
			result, err := SignupHandler(ctx)
			if err != nil {
				structure.Fail(ctx, err.Status, err.Message)
				return
			}
			structure.OK(ctx, http.StatusCreated, result, "user created successfully")
		},
	)

	AuthenticationRoute.POST(
		"/login",
		pipe.Validate(func() any {
			return &dto.Login{}
		}),
		func(ctx *gin.Context) {
			result, err := LoginHandler(ctx)
			if err != nil {
				structure.Fail(ctx, err.Status, err.Message)
				return
			}
			structure.OK(ctx, http.StatusOK, result, "login successfully")
		},
	)
}
