package dependencies

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/controller"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
)

func ProductDependency() controller.IProductController {
	productRepository := repository.NewProductRepository(database.Get())
	productUseCases := usecases.NewProductUseCases(productRepository)
	productController := controller.NewProductController(productUseCases)

	return productController
}
