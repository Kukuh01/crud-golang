package repositories

import (
	"github.com/kukuh01/crud-golang/config"
	"github.com/kukuh01/crud-golang/models"
)

func GetAllProducts(products *[]models.Product) error {
	return config.DB.Find(products).Error
}

func GetProductByID(product *models.Product, id uint) error {
	return config.DB.First(product, id).Error
}

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(product *models.Product) error {
	return config.DB.Delete(product).Error
}
