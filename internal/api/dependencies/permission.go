package dependencies

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/controller"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
)

func PermissionDependency() controller.IPermissionController {
	permissionRepository := repository.NewPermissionRepository(database.Get())
	permissionUseCases := usecases.NewPermissionUseCases(permissionRepository)
	permissionController := controller.NewPermissionController(permissionUseCases)

	return permissionController
}
