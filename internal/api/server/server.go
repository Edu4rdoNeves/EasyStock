package server

import (
	"net/http"
	"os"

	"github.com/Edu4rdoNeves/EasyStrock/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	server := gin.Default()
	router := router.Router(server)

	router.GET("/health", healthRoute)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

func healthRoute(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Health Response Okay")
}
