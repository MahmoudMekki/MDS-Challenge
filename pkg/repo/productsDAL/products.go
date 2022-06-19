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

func GetProdBySkuAndCountry(sku, country string) (prod models.Product, existed bool, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return prod, existed, err
	}
	err = dbConn.Table(models.ProductsTableName).Where("sku=? and country_code=?", sku, country).Find(&prod).Error
	if err != nil {
		return prod, false, err
	}
	if prod.Id <= 0 {
		return prod, false, nil
	}
	return prod, true, nil
}

func UpdateProdStock(productID int, change int) error {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return err
	}
	err = dbConn.Exec("update products set amount = amount+? where id=?", change, productID).Error
	return err
}
