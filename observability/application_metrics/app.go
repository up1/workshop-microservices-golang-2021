package demo

import (
	"demo/beer"
	_ "demo/docs"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func StartServer() {
	r := gin.New()
	r.Use(gin.Recovery())

	// ===== Prometheus
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	// set middleware for gin
	m.Use(r)

	// Docs
	// r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add routing
	beer.NewRoutes(r)

	r.Run()
}
