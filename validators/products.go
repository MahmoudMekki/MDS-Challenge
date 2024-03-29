package validators

import (
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ValidateCreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		data, err := jio.ValidateJSON(&jsonData, jio.Object().Keys(jio.K{
			"sku":          jio.String().Min(3).Required(),
			"name":         jio.String().Min(3).Required(),
			"country_code": jio.String().Required(),
			"amount":       jio.Number().Required(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("sku", data["sku"])
		ctx.Set("name", data["name"])
		ctx.Set("country_code", data["country_code"])
		ctx.Set("amount", data["amount"])
		ctx.Next()
	}
}
func ValidateUpdateProduct() gin.HandlerFunc {
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
			"name":         jio.String().Min(3).Optional(),
			"country_code": jio.String().Optional(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("sku", sku)
		ctx.Set("name", data["name"])
		ctx.Set("country_code", data["country_code"])
		ctx.Next()
	}
}
func ValidateGetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sku, ok := ctx.Params.Get("sku")
		if !ok || len(sku) < 3 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid sku"})
			return
		}
		ctx.Set("sku", sku)
		ctx.Next()
	}
}
func ValidateGetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyword, _ := ctx.GetQuery("keyword")
		limitStr, _ := ctx.GetQuery("limit")
		pageStr, _ := ctx.GetQuery("page")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		ctx.Set("limit", limit)
		ctx.Set("page", page)
		ctx.Set("keyword", keyword)
		ctx.Next()
	}
}
