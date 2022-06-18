package product

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func UpdateProduct(ctx *gin.Context) {
	prod := models.Product{
		SKU:         ctx.GetString("sku"),
		CountryCode: ctx.GetString("country_code"),
		Name:        ctx.GetString("name"),
	}
	prod, err := productsDAL.UpdateProduct(prod)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting the prod"})
		return
	}
	ctx.JSON(http.StatusOK, prod)
}
