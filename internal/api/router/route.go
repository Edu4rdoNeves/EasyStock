package router

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/api/dependencies"
	"github.com/Edu4rdoNeves/EasyStrock/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	loginControllerWithDependencies := dependencies.LoginDependency()
	userControllerWithDependencies := dependencies.UserDependency()
	productControllerWithDependencies := dependencies.ProductDependency()
	permissionControllerWithDependencies := dependencies.PermissionDependency()

	main := router.Group("api/v1")
	{
		login := main.Group("login")
		{
			login.POST("/", loginControllerWithDependencies.Login)
		}

		user := main.Group("user", middleware.Auth())
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
			product.GET("/:id", productControllerWithDependencies.GetProductById)
			product.PUT("/:id", productControllerWithDependencies.UpdateProduct)
			product.DELETE("/:id", productControllerWithDependencies.DeleteProduct)
			product.POST("/", productControllerWithDependencies.CreateProduct)
		}

		permission := main.Group("permission", middleware.Auth())
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
