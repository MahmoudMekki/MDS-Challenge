package product

import (
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func GetProducts(ctx *gin.Context) {
	paginator := models.Paginator{
		Limit:   ctx.GetInt("limit"),
		Page:    ctx.GetInt("page"),
		KeyWord: ctx.GetString("keyword"),
	}
	prods, hits, err := productsDAL.GetProducts(paginator)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting prods"})
		return
	}
	ctx.Header("X-TOTAL-PRODUCTS", strconv.Itoa(int(hits)))
	ctx.JSON(http.StatusOK, prods)
}
