package router

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/api/dependencies"
	"github.com/Edu4rdoNeves/EasyStrock/internal/api/middleware"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	loginControllerWithDependencies := dependencies.LoginDependency()
	userControllerWithDependencies := dependencies.UserDependency()
	productControllerWithDependencies := dependencies.ProductDependency()
	permissionControllerWithDependencies := dependencies.PermissionDependency()

	db := database.Get()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(204)
		return
	})

	main := router.Group("api/v1")
	{
		login := main.Group("login")
		{
			login.POST("/", loginControllerWithDependencies.Login)
		}

		user := main.Group("user", middleware.Auth(), middleware.AdminMiddleware(db))
		{
			user.GET("/", userControllerWithDependencies.GetUsers)
			user.GET("/:id", userControllerWithDependencies.GetUserById)
			user.PUT("/:id", userControllerWithDependencies.UpdateUser)
			user.DELETE("/:id", userControllerWithDependencies.DeleteUser)
		}

		create := main.Group("user/create")
		{
			create.POST("/", userControllerWithDependencies.CreateUser)
		}

		product := main.Group("product", middleware.Auth())
		{
			product.GET("/", productControllerWithDependencies.GetProducts)
			product.PUT("/:id", productControllerWithDependencies.UpdateProduct)
			product.DELETE("/:id", productControllerWithDependencies.DeleteProduct)
			product.POST("/", productControllerWithDependencies.CreateProduct)
		}

		permission := main.Group("permission", middleware.Auth(), middleware.AdminMiddleware(db))
		{
			permission.GET("/", permissionControllerWithDependencies.GetPermissions)
			permission.GET("/:id", permissionControllerWithDependencies.GetPermissionById)
			permission.PUT("/:id", permissionControllerWithDependencies.UpdatePermission)
			permission.DELETE("/:id", permissionControllerWithDependencies.DeletePermission)
			permission.POST("/", permissionControllerWithDependencies.CreatePermission)
		}
	}
	return router
}
