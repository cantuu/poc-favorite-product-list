package controllers

import (
	"fmt"
	"log"
	"product/config"
	"product/models"

	// "product/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite database driver
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize() {

	var err error

	DBURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.ENV.DBHost,
		config.ENV.DBPort,
		config.ENV.DBUser,
		config.ENV.DBName,
		config.ENV.DBPassword,
	)
	server.DB, err = gorm.Open(config.ENV.DBDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", config.ENV.DBDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", config.ENV.DBDriver)
	}

	server.DB.Debug().AutoMigrate(&models.Product{}) //database migration

	server.Router = server.SetupRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port: ", config.ENV.ApiPort)

	server.Router.Run(fmt.Sprint(":", config.ENV.ApiPort))
}
