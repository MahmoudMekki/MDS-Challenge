package router

import (
	"github.com/MahmoudMekki/MDS-task/cmd/product"
	"github.com/MahmoudMekki/MDS-task/validators"
)

func (r *routerImp) setProductRoutes() {
	productEndpoints := r.engine.Group("/api/v1/products")
	productEndpoints.GET("", validators.ValidateGetProducts(), product.GetProducts)
	productEndpoints.GET("/:sku", validators.ValidateGetProduct(), product.GetProduct)
	productEndpoints.POST("", validators.ValidateCreateProduct(), product.CreateProduct)
	productEndpoints.PUT("/:sku", validators.ValidateUpdateProduct(), product.UpdateProduct)
}
