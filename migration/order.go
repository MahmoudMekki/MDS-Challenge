package migration

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listOrder() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-orders-table-and-id-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Order{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.OrdersTableName)
			},
		},
	}
}
