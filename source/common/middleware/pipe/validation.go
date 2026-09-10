package pipe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(bodyFactory func() any) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := bodyFactory()

		if err := ctx.ShouldBindJSON(body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.Set("body", body)
		ctx.Next()
	}
}