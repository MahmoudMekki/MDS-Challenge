package migration

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listProduct() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-products-table-and-id-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Product{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.ProductsTableName)
			},
		},
	}
}
