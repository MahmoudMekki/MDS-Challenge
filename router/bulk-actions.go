package router

import (
	"github.com/MahmoudMekki/MDS-task/cmd/orders"
)

func (r *routerImp) setBulkRoutes() {
	productEndpoints := r.engine.Group("/api/v1/bulk")
	productEndpoints.POST("/orders", orders.BulkUpload)
}
