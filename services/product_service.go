package services

import (
	"github.com/kukuh01/crud-golang/models"
	"github.com/kukuh01/crud-golang/repositories"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := repositories.GetAllProducts(&products)
	return products, err
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := repositories.GetProductByID(&product, id)
	return product, err
}

func CreateProduct(product models.Product) error {
	return repositories.CreateProduct(&product)
}

func UpdateProduct(product models.Product) error {
	return repositories.UpdateProduct(&product)
}

func DeleteProduct(product models.Product) error {
	return repositories.DeleteProduct(&product)
}
