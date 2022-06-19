package orderDAL

import (
	"github.com/MahmoudMekki/MDS-task/database"
	"github.com/MahmoudMekki/MDS-task/pkg/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Order{}, err
	}
	err = dbConn.Table(models.OrdersTableName).Create(&order).Error
	return order, err
}
