package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 상품 목록 조회
func (c *Controller) ReadProducts(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	products, err := c.service.ReadProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, products)
}

// 프로모션 상품 목록 조회
func (c *Controller) ReadPromotedProducts(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	promotedProducts, err := c.service.ReadPromotedProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, promotedProducts)
}
