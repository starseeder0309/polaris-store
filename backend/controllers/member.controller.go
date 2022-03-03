package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kimhyeonu/polaris-store/backend/models"
	"github.com/kimhyeonu/polaris-store/backend/services"
)

// 회원 생성
func (c *Controller) CreateMember(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var newMember models.Member

	err := context.ShouldBindJSON(&newMember)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMember, err = c.service.CreateMember(newMember)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, newMember)
}

// 회원 접속
func (c *Controller) ConnectMember(context *gin.Context) {
	if c.service == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var member models.Member

	err := context.ShouldBindJSON(&member)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err = c.service.ConnectMember(member.Email, member.Password)
	if err != nil {
		if err == services.ErrInvalidPassword {
			context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, member)
}

// 회원 접속 해제
func (c *Controller) DisconnectMember(context *gin.Context) {
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

	err = c.service.DisconnectMember(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
