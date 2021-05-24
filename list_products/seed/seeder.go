package seed

import (
	"list_product/models"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Name:     "Gabriel Cantu",
		Email:    "gabrielcantu11@gmail.com",
		Password: "12345",
	},
	{
		Name:     "Batatinha da Silva",
		Email:    "batatinha@batatas.com",
		Password: "09876",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
