package controller

import (
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type IPermissionController interface {
	GetPermissions(context *gin.Context)
	GetPermissionById(context *gin.Context)
	CreatePermission(context *gin.Context)
	UpdatePermission(context *gin.Context)
	DeletePermission(context *gin.Context)
}

type PermissionController struct {
	usecases usecases.IPermissionUseCases
}

func NewPermissionController(usecases usecases.IPermissionUseCases) IPermissionController {
	return &PermissionController{usecases}
}

func (c *PermissionController) GetPermissions(context *gin.Context) {

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	PermissionResponse, err := c.usecases.GetPermissions(page, limit)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't find Permissions: " + err.Error(),
		})
		return
	}
	context.JSON(200, PermissionResponse)
}

func (c *PermissionController) GetPermissionById(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	Permission, err := c.usecases.GetPermissionById(param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't find a Permission: " + err.Error(),
		})
		return
	}

	context.JSON(200, Permission)
}

func (c *PermissionController) CreatePermission(context *gin.Context) {
	Permission := &model.Permission{}

	err := context.ShouldBindJSON(Permission)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.CreatePermission(Permission)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't create Permission: " + err.Error(),
		})
		return
	}

	context.JSON(201, gin.H{
		"Permission Created:": Permission,
	})
}

func (c *PermissionController) UpdatePermission(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	PermissionUpdate := &model.Permission{}
	err := context.ShouldBindJSON(&PermissionUpdate)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.UpdatePermission(PermissionUpdate, param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't update a Permission: " + err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"Permission Updated:": PermissionUpdate,
	})
}

func (c *PermissionController) DeletePermission(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	err := c.usecases.DeletePermission(param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't delete a Permission: " + err.Error(),
		})
		return
	}

	context.JSON(200, "Permission deleted")
}
