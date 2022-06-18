package productsDAL

import (
	"fmt"
	"github.com/MahmoudMekki/MDS-task/database"
	"github.com/MahmoudMekki/MDS-task/pkg/models"
)

func CreateProduct(prod models.Product) (models.Product, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Product{}, err
	}
	err = dbConn.Table(models.ProductsTableName).Create(&prod).Error
	return prod, err
}

func UpdateProduct(prod models.Product) (models.Product, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Product{}, err
	}
	err = dbConn.Table(models.ProductsTableName).Where("sku=?", prod.SKU).Updates(&prod).Error
	return prod, err
}

func GetProduct(sku string) (prods []models.Product, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return nil, err
	}
	err = dbConn.Preload("Orders").Table(models.ProductsTableName).Where("sku=?", sku).Find(&prods).Error
	return prods, err
}
func GetProducts(paginator models.Paginator) (prods []models.Product, hits int64, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return nil, 0, err
	}
	err = dbConn.Table(models.ProductsTableName).
		Where(" name LIKE ?", fmt.Sprintf("%s%s%s", "%", paginator.KeyWord, "%")).
		Count(&hits).
		Offset(paginator.GetOffset()).Limit(paginator.GetLimit()).Find(&prods).Error
	return prods, hits, err
}
