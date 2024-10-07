package dependencies

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/controller"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
)

func UserDependency() controller.IUserController {
	userRepository := repository.NewUserRepository(database.Get())
	userUseCases := usecases.NewUserUseCases(userRepository)
	userController := controller.NewUserController(userUseCases)

	return userController
}
