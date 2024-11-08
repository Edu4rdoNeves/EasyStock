package controller

import (
	"net/http"
	"strconv"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/usecases"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type IProductController interface {
	GetProducts(context *gin.Context)
	CreateProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
}

type ProductController struct {
	usecases usecases.IProductUseCases
}

func NewProductController(usecases usecases.IProductUseCases) IProductController {
	return &ProductController{usecases}
}

func (c *ProductController) GetProducts(context *gin.Context) {

	param := context.Query("param")

	if param != "" {
		product, err := c.usecases.GetProductByNameOrID(param)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"message": "Produto não encontrado"})
			return
		}

		products := []model.Product{
			*product,
		}

		context.JSON(200, products)
		return
	}

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	products, err := c.usecases.GetProducts(page, limit)
	if err != nil {
		context.JSON(500, gin.H{
			"Error:": "Can't get products" + err.Error(),
		})
		return
	}

	context.JSON(200, products)
}

func (c *ProductController) CreateProduct(context *gin.Context) {
	product := &model.Product{}

	err := context.ShouldBindJSON(product)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.CreateProduct(product)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't create product: " + err.Error(),
		})
		return
	}

	context.JSON(201, gin.H{
		"Product Created:": product,
	})
}

func (c *ProductController) UpdateProduct(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	productUpdate := &model.Product{}
	err := context.ShouldBindJSON(&productUpdate)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON: " + err.Error(),
		})
		return
	}

	err = c.usecases.UpdateProduct(productUpdate, param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't update a product: " + err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"Product Updated:": productUpdate,
	})
}

func (c *ProductController) DeleteProduct(context *gin.Context) {
	param, bool := context.Params.Get("id")
	if !bool {
		context.JSON(500, gin.H{
			"Error: ": "Param is not valid",
		})
		return
	}

	err := c.usecases.DeleteProduct(param)
	if err != nil {
		context.JSON(500, gin.H{
			"Error: ": "Can't delete a product: " + err.Error(),
		})
		return
	}

	context.JSON(200, "Product deleted")
}
