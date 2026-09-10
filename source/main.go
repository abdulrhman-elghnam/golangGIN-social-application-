package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"golang/source/app"
	"golang/source/common/middleware/global"
	"golang/source/common/structure"
	"golang/source/database"
	"golang/source/module/authentication"
	"golang/source/module/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connection()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(global.ErrorHandler())
	bootstrap := router.Group("/")


	app.RegisterMainRoutes(bootstrap)
	authentication.RegisterAuthenticationRoutes(bootstrap)
	user.RegisterUserRoutes(bootstrap)

	router.NoRoute(func(c *gin.Context) {
		structure.Fail(
			c,
			http.StatusNotFound,
			"NOT_FOUND",
			"route not found ❌",
		)
	})

	router.Run(":" + os.Getenv("PORT"))
}
