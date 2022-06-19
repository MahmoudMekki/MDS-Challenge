package router

import (
	"github.com/MahmoudMekki/MDS-task/cmd/orders"
	"github.com/MahmoudMekki/MDS-task/validators"
)

func (r *routerImp) setOrderRoutes() {
	productEndpoints := r.engine.Group("/api/v1/products/:sku/orders")
	productEndpoints.POST("", validators.ValidateCreateOrder(), orders.CreateOrder)
}
