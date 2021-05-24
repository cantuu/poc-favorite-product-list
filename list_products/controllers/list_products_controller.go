package controllers

import (
	"fmt"
	"list_product/models"
	"list_product/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllListOfProducts(context *gin.Context) {

	listModel := models.ListProducts{}

	lists, err := listModel.GetAllListsOfProducts(server.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"mensagem": "Houve um erro ao listas todos os usuários",
		})
		return
	}
	context.JSON(http.StatusOK, lists)

}

func (server *Server) DeleteListOfProduct(context *gin.Context) {
	paramId, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}

	list := models.ListProducts{}
	_, err = list.DeleteListOfProduct(server.DB, uint32(paramId))
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusNoContent, "")
}

func (server *Server) GetListOfProducts(context *gin.Context) {
	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}
	list := models.ListProducts{}
	listGotten, err := list.FindListByID(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, listGotten)
}

func (server *Server) AddProductOnList(context *gin.Context) {
	list := models.ListProducts{}

	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}

	listGotten, err := list.FindListByID(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var listRequest models.ListProductRequest
	if err := context.BindJSON(&listRequest); err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	productRepo := repository.NewProductRepository()
	var productsToAdd []int64 = listGotten.Items

	for _, item := range listRequest.Itens {
		if IsItemInList(item, listGotten.Items) {
			fmt.Println("Item ja esta na lista")
			continue
		}

		product, err := productRepo.GetDataFromProductMS(item)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		productsToAdd = append(productsToAdd, int64(product.ID))
	}

	listGotten.Items = productsToAdd
	listProductCreated, err := listGotten.UpdateList(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, listProductCreated)
}

func (server *Server) DeleteProductFromList(context *gin.Context) {
	list := models.ListProducts{}

	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}

	listGotten, err := list.FindListByID(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var listRequest models.ListProductRequest
	if err := context.BindJSON(&listRequest); err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	productRepo := repository.NewProductRepository()
	var productsToRemove []int64 = listGotten.Items

	for i, item := range listRequest.Itens {
		if !IsItemInList(item, listGotten.Items) {
			fmt.Println("Item nao esta na lista")
			continue
		}

		_, err := productRepo.GetDataFromProductMS(item)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		productsToRemove = RemoveProductFromList(productsToRemove, i)
	}
	fmt.Println(productsToRemove)
	listGotten.Items = productsToRemove

	listProductCreated, err := listGotten.UpdateList(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, listProductCreated)
}

func IsItemInList(item int, list []int64) bool {
	for _, v := range list {
		if item == int(v) {
			return true

		}
	}
	return false
}

func RemoveProductFromList(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
