package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 회원 ID로 주문 목록 조회
func (c *Controller) ReadOrdersById(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := c.service.ReadOrdersById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, orders)
}
