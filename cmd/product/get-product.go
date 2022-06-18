package product

import (
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetProduct(ctx *gin.Context) {
	sku := ctx.GetString("sku")
	prods, err := productsDAL.GetProduct(sku)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting the prod"})
		return
	}
	if len(prods) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no prod with this SKU"})
		return
	}
	ctx.JSON(http.StatusOK, prods)
}
