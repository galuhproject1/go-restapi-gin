package main

import (
	productcontroller "github.com/galuhproject1/go-restapi-gin/controllers/productController"
	"github.com/galuhproject1/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)
	r.Run()
}
