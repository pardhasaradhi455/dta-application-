package controller

import "github.com/gin-gonic/gin"

type OrderControllerInterface interface {
	GetOrder(c *gin.Context)
	GetStatus(c *gin.Context)
	GetAllOrders(c *gin.Context)
}
