package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Price       string    `gorm:"size:255;not null;unique" json:"price"`
	Image       string    `gorm:"size:100;not null;" json:"image"`
	Brand       string    `gorm:"size:100;not null;" json:"brand"`
	Title       string    `gorm:"size:100;not null;" json:"title"`
	ReviewScore string    `gorm:"size:100;not null;" json:"reviewScore"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (product *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	users := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &users, err
}

func (product *Product) FindProductByID(db *gorm.DB, uid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&product).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}
	return product, err
}

func (product *Product) CreateProduct(db *gorm.DB) (*Product, error) {

	var err error
	err = db.Debug().Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (product *Product) UpdateProduct(db *gorm.DB, uid uint32) (*Product, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"title":       product.Title,
			"brand":       product.Brand,
			"image":       product.Image,
			"price":       product.Price,
			"reviewScore": product.ReviewScore,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &Product{}, db.Error
	}

	err := db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (product *Product) DeleteProduct(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (product *Product) Prepare() {
	product.ID = 0
	product.Brand = html.EscapeString(strings.TrimSpace(product.Brand))
	product.Image = html.EscapeString(strings.TrimSpace(product.Image))
	product.Price = html.EscapeString(strings.TrimSpace(product.Price))
	product.ReviewScore = html.EscapeString(strings.TrimSpace(product.ReviewScore))
	product.Title = html.EscapeString(strings.TrimSpace(product.Title))

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
}
