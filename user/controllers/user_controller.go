package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user/auth"
	"user/models"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetAllUsers(context *gin.Context) {

	userModel := models.User{}

	users, err := userModel.GetAllUsers(server.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"mensagem": "Houve um erro ao listas todos os usuários",
		})
		return
	}
	context.JSON(http.StatusOK, users)

}

func (server *Server) GetUser(context *gin.Context) {
	uid, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, userGotten)
}

func (server *Server) CreateUser(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		return // throws Bad Request
	}

	err := user.PrepareAndValidate()
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	sqsMessage, _ := json.Marshal(userCreated)

	err = server.SQSMessager.SendMessage(string(sqsMessage))
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, userCreated)
}

func (server *Server) UpdateUser(context *gin.Context) {
	paramId, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}

	var user models.User
	if err := context.BindJSON(&user); err != nil {
		return // throws Bad Request
	}

	uid, err := auth.ExtractTokenID(context)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if uid != uint32(paramId) {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Usuário não autorizado a atualizar outro uid",
		})
		return
	}

	err = user.PrepareAndValidate()
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	updatedUser, err := user.UpdateUser(server.DB, uint32(paramId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}
	context.JSON(http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(context *gin.Context) {
	paramId, err := strconv.ParseUint(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Houve um erro ao obter o ID da requisição",
		})
		return
	}

	uid, err := auth.ExtractTokenID(context)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if uid != uint32(paramId) {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Usuário não autorizado a atualizar outro uid",
		})
		return
	}

	user := models.User{}
	_, err = user.DeleteUser(server.DB, uint32(uid))
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusNoContent, "")
}
