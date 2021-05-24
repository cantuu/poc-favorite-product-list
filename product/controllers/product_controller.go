package controllers

import (
	"net/http"
	"product/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllProducts(context *gin.Context) {

	productModel := models.Product{}

	products, err := productModel.FindAllProducts(server.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Houve um erro ao listas todos os produtos"})
		return
	}
	context.JSON(http.StatusOK, products)

}

func (server *Server) GetProduct(context *gin.Context) {
	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Houve um erro ao obter o ID da requisição"})
		return
	}
	product := models.Product{}
	productGotten, err := product.FindProductByID(server.DB, uint32(uid))
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, productGotten)
}

func (server *Server) CreateProduct(context *gin.Context) {
	var product models.Product
	if err := context.BindJSON(&product); err != nil {
		return // throws Bad Request
	}

	product.Prepare()

	productCreated, err := product.CreateProduct(server.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, productCreated)
}

func (server *Server) UpdateProduct(context *gin.Context) {
	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Houve um erro ao obter o ID da requisição"})
		return
	}

	var product models.Product
	if err := context.BindJSON(&product); err != nil {
		return // throws Bad Request
	}

	product.Prepare()

	updatedProduct, err := product.UpdateProduct(server.DB, uint32(uid))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, updatedProduct)
}

func (server *Server) DeleteProduct(context *gin.Context) {
	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Houve um erro ao obter o ID da requisição"})
		return
	}

	product := models.Product{}

	_, err = product.DeleteProduct(server.DB, uint32(uid))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusNoContent, "")
}
