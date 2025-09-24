package migrations

import (
	"github.com/kukuh01/crud-golang/config"
	"github.com/kukuh01/crud-golang/models"
)

func RunMigration() {
	config.DB.AutoMigrate(&models.Product{})
}
