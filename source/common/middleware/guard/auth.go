package guard

import (
    "fmt"
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

        parts := strings.Fields(authHeader)

        if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
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

func UserID(c *gin.Context) (uint, bool) {
    claims, exists := c.Get("claims")
    if !exists {
        return 0, false
    }

    value, ok := claims.(map[string]interface{})["user_id"]
    if !ok {
        return 0, false
    }

    var id uint
    switch typed := value.(type) {
    case float64:
        id = uint(typed)
    case float32:
        id = uint(typed)
    case int:
        id = uint(typed)
    case uint:
        id = typed
    default:
        return 0, false
    }

    if id == 0 {
        return 0, false
    }
    return id, true
}

var _ = fmt.Sprint