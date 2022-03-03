package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kimhyeonu/polaris-store/backend/services"
)

type Controller struct {
	service services.Service
}

func GenerateController() (*Controller, error) {
	return new(Controller), nil
}

func (c *Controller) ProcessPayment(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
}
