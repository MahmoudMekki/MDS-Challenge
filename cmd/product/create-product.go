package product

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateProduct(ctx *gin.Context) {
	product := models.Product{
		SKU:         ctx.GetString("sku"),
		CountryCode: ctx.GetString("country_code"),
		Name:        ctx.GetString("name"),
		Amount:      int(ctx.GetFloat64("amount")),
	}
	product, err := productsDAL.CreateProduct(product)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "this product is already existed"})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}
