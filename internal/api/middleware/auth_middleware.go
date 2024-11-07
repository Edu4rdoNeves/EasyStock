package middleware

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/tools"
	"github.com/gin-gonic/gin"
)

const admin = 1

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Permite que requisições OPTIONS passem sem autenticação
		if context.Request.Method == "OPTIONS" {
			context.Next()
			return
		}

		const Bearer_schema = "Bearer "
		header := context.GetHeader("Authorization")
		if header == "" {
			context.AbortWithStatus(401)
			return
		}

		token := header[len(Bearer_schema):]

		if !tools.NewJWTService().ValidateToken(token) {
			context.AbortWithStatus(401)
			return
		}

		context.Next() // Continue para a próxima etapa, se autenticado
	}
}
