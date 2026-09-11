package app

import (
	"golang/source/common/structure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainController(ctx *gin.Context) {
	structure.OK(ctx, http.StatusOK ,  map[string]any{
		"msg": "hi from backend server 🚀",
	}, "")
}
