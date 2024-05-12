package main

import (
	productcontroller "github.com/galuhproject1/go-restapi-gin/controllers/productController"
	"github.com/galuhproject1/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/products", productcontroller.Index)
	r.GET("/products/:id", productcontroller.Show)
	r.POST("/products", productcontroller.Create)
	r.PUT("/products/:id", productcontroller.Update)
	r.DELETE("/products/:id", productcontroller.Delete)
	r.Run()
}
