package user

import (
	"golang/source/common/structure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	structure.OK(c, http.StatusOK, gin.H{
		"message": "welcome from backend server 🚀",
	})
}

