package seed

import (
	"log"
	"product/models"

	"github.com/jinzhu/gorm"
)

var products = []models.Product{
	{
		Price:       "1234",
		Title:       "Carrinho",
		ReviewScore: "23445",
		Brand:       "sdfasdfasdf",
		Image:       "http://asdasdasd.com.br/carrinho",
	},
	{
		Price:       "5678",
		Title:       "Boneca",
		ReviewScore: "3241234",
		Brand:       "1234234214",
		Image:       "http://asdasdasd.com.br/boneca",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i := range products {
		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
