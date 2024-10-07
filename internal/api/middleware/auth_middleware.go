package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/Edu4rdoNeves/EasyStrock/internal/tools"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const admin = 1

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		const Bearer_schena = "Bearer "
		header := context.GetHeader("Authorization")
		if header == "" {
			context.AbortWithStatus(401)
		}

		token := header[len(Bearer_schena):]

		if !tools.NewJWTService().ValidateToken(token) {
			context.AbortWithStatus(401)
			return
		}
	}
}

func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &tools.Claim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*tools.Claim)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := claims.Sum
		var user = &model.Users{}

		err = db.First(&user, userID).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if user.PermissionID != admin {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
