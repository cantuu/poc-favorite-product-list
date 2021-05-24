package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"list_product/config"
	"list_product/models"
	"net/http"
	"strconv"
	"time"
)

type ProductRepository struct{}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (*ProductRepository) GetDataFromProductMS(productId int) (*models.Product, error) {
	url := config.ENV.ProductEndpoint + "/api/product/GetById/" + strconv.Itoa(productId)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	product := models.Product{}
	err = json.Unmarshal(data, &product)
	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, errors.New("Product not found")
	}

	return &product, nil

}
