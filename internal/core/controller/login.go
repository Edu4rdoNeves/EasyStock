package controller

import (
	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	Login(context *gin.Context)
}

type LoginController struct {
	business usecases.ILoginUseCases
}

func NewLoginController(iUseCases usecases.ILoginUseCases) ILoginController {
	return &LoginController{iUseCases}
}

func (c *LoginController) Login(context *gin.Context) {
	login := &model.Login{}

	err := context.ShouldBindBodyWithJSON(login)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
		return
	}

	token, err := c.business.Login(login)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"token": token,
	})
}
