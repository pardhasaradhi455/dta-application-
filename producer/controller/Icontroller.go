package controller

import "github.com/gin-gonic/gin"

type OrderControllerInterface interface{
	PlaceOrder(c *gin.Context)
	GetPendingOrders(c *gin.Context)
	GetDeliveredOrders(c *gin.Context)
	ChangeState(c *gin.Context)
	ChangeAddress(c *gin.Context)
}