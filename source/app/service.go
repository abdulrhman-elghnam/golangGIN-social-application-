package app

import (
	"golang/source/common/structure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainController(c *gin.Context) {
	structure.OK(c, http.StatusOK, gin.H{
		"message": "welcome from backend server 🚀",
	})
}
