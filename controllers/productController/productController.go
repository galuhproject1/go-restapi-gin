package productcontroller

import (
	"net/http"

	"github.com/galuhproject1/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
