package controller

import (
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/Edu4rdoNeves/EasyStrock/internal/tools"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUsers(context *gin.Context)
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
}

type UserController struct {
	usecases usecases.IUserUseCases
}

func NewUserController(usecases usecases.IUserUseCases) IUserController {
	return &UserController{usecases}
}

func (c *UserController) GetUsers(context *gin.Context) {

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	userResponse, err := c.usecases.GetUsers(page, limit)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't find users: " + err.Error(),
		})
		return
	}
	context.JSON(200, userResponse)
}

func (c *UserController) GetUserById(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	user, err := c.usecases.GetUserById(param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't find a user: " + err.Error(),
		})
		return
	}

	context.JSON(200, user)
}

func (c *UserController) CreateUser(context *gin.Context) {
	user := &model.Users{}

	err := context.ShouldBindJSON(user)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.CreateUser(user)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't create user: " + err.Error(),
		})
		return
	}

	user.Password = tools.SHA256Enconder(user.Password)

	context.JSON(201, gin.H{
		"User Created:": user,
	})
}

func (c *UserController) UpdateUser(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	userUpdate := &model.Users{}
	err := context.ShouldBindJSON(&userUpdate)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.UpdateUser(userUpdate, param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't update a user: " + err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"User Updated:": userUpdate,
	})
}

func (c *UserController) DeleteUser(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	err := c.usecases.DeleteUser(param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't delete a user: " + err.Error(),
		})
		return
	}

	context.JSON(200, "User deleted")
}
