package controllers

import (
	"net/http"
	"user/auth"
	"user/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(context *gin.Context) {
	user := models.User{}
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

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (server *Server) SignIn(email, password string) (string, error) {

	user := models.User{}

	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
