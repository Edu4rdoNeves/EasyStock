package dependencies

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/controller"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
)

func LoginDependency() controller.ILoginController {
	loginRepository := repository.NewLoginRepository(database.Get())
	loginUseCases := usecases.NewLoginBusiness(loginRepository)
	loginController := controller.NewLoginController(loginUseCases)

	return loginController
}
