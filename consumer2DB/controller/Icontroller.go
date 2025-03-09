package controller

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	FetchOrderByKey(c *gin.Context)
	InsertOrder(c *gin.Context)
	FetchAllOrders(c *gin.Context)
}