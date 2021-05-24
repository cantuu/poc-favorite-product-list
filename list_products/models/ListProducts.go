package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type ListProducts struct {
	ID        uint32        `gorm:"auto_increment" json:"id"`
	UserId    uint32        `gorm:"primary_key" json:"user_id"`
	Items     pq.Int64Array `gorm:"type:integer[]" json:"items"`
	CreatedAt time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Items struct {
	gorm.Model
	Item int
}

type ListProductRequest struct {
	Itens []int `json:"items" binding:"required"`
}

func (listProducts *ListProducts) SaveListProduct(db *gorm.DB) (*ListProducts, error) {
	var err error
	err = db.Debug().Create(&listProducts).Error
	if err != nil {
		return &ListProducts{}, err
	}
	return listProducts, nil
}

func (listProducts *ListProducts) GetAllListsOfProducts(db *gorm.DB) (*[]ListProducts, error) {
	var err error
	lists := []ListProducts{}
	err = db.Debug().Model(&ListProducts{}).Limit(100).Find(&lists).Error
	if err != nil {
		return &[]ListProducts{}, err
	}
	return &lists, err
}

func (listProducts *ListProducts) DeleteListOfProduct(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&ListProducts{}).Where("id = ?", uid).Take(&ListProducts{}).Delete(&ListProducts{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (listProducts *ListProducts) FindListByID(db *gorm.DB, uid uint32) (*ListProducts, error) {
	var err error
	err = db.Debug().Model(ListProducts{}).Where("user_id = ?", uid).Take(&listProducts).Error
	if err != nil {
		return &ListProducts{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &ListProducts{}, errors.New("ListProducts Not Found")
	}
	return listProducts, err
}

// func Hash(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func (listProducts *ListProducts) BeforeSave() error {
// 	hashedPassword, err := Hash(listProducts.Password)
// 	if err != nil {
// 		return err
// 	}
// 	listProducts.Password = string(hashedPassword)
// 	return nil
// }

// func (listProducts *ListProducts) PrepareAndValidate() error {
// 	listProducts.ID = 0
// 	listProducts.Name = html.EscapeString(strings.TrimSpace(listProducts.Name))
// 	listProducts.Email = html.EscapeString(strings.TrimSpace(listProducts.Email))
// 	listProducts.CreatedAt = time.Now()
// 	listProducts.UpdatedAt = time.Now()

// 	if listProducts.Name == "" {
// 		return errors.New("Required Name")
// 	}
// 	if listProducts.Password == "" {
// 		return errors.New("Required Password")
// 	}
// 	if listProducts.Email == "" {
// 		return errors.New("Required Email")
// 	}
// 	if err := checkmail.ValidateFormat(listProducts.Email); err != nil {
// 		return errors.New("Invalid Email")
// 	}

// 	return nil
// }

func (listProducts *ListProducts) UpdateList(db *gorm.DB, uid uint32) (*ListProducts, error) {

	db = db.Debug().Model(&ListProducts{}).Where("user_id = ?", uid).Take(&ListProducts{}).UpdateColumns(
		map[string]interface{}{
			"id":         listProducts.ID,
			"items":      listProducts.Items,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &ListProducts{}, db.Error
	}

	err := db.Debug().Model(&ListProducts{}).Where("user_id = ?", uid).Take(&listProducts).Error
	if err != nil {
		return &ListProducts{}, err
	}
	return listProducts, nil
}
