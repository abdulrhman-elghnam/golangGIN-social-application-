package guard

import (
    "net/http"
    "os"
    "strings"

    "golang/source/common/jwt"

    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {

        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Authorization header is required",
            })
            return
        }

        parts := strings.Split(authHeader, " ")

        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Invalid authorization format",
            })
            return
        }

        token := parts[1]
        secret := os.Getenv("TKN_KEY")

        claims, err := jwt.ValidateToken(token, secret)

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Invalid or expired token",
            })
            return
        }

        c.Set("claims", claims)

        c.Next()
    }
}