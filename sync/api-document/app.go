package demo

import (
	"demo/beer"
	_ "demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	r := gin.New()
	r.Use(gin.Recovery())

	// Docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add routing
	beer.NewRoutes(r)

	r.Run()
}
