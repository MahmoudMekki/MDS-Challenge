package orders

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/orderDAL"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateOrder(ctx *gin.Context) {
	sku := ctx.GetString("sku")
	country := ctx.GetString("country_code")
	amount := int(ctx.GetFloat64("amount"))
	prod, existed, err := productsDAL.GetProdBySkuAndCountry(sku, country)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the order"})
		return
	}
	if !existed {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no product with  this sku"})
		return
	}
	if prod.Amount+amount < 0 {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "product stock is insufficient"})
		return
	}
	order := models.Order{
		ProductId: prod.Id,
		Amount:    amount,
	}
	err = productsDAL.UpdateProdStock(prod.Id, amount)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the order"})
		return
	}
	order, err = orderDAL.CreateOrder(order)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the order"})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}
