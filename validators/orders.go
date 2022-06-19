package validators

import (
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ValidateCreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sku, ok := ctx.Params.Get("sku")
		if !ok || len(sku) < 3 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid sku"})
			return
		}
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		data, err := jio.ValidateJSON(&jsonData, jio.Object().Keys(jio.K{
			"amount":       jio.Number().Required(),
			"country_code": jio.String().Required(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("sku", sku)
		ctx.Set("amount", data["amount"])
		ctx.Set("country_code", data["country_code"])
		ctx.Next()
	}
}
