package controller

import (
	"delivery_tracking_api/consumer1DB/logger"
	"delivery_tracking_api/consumer1DB/model"
	"delivery_tracking_api/consumer1DB/repo"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	ControllerInterface
}

var Repo = &repo.Repo{}

func (ctrl *OrderController) InsertOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	logger.LogInfo("Invoked controller to insert order with key" + order.OrderId, c)
	c.JSON(http.StatusCreated, Repo.InserOrUpdateRecord(order.OrderId, order))
}

func (ctrl *OrderController) FetchOrderByKey(c *gin.Context) {
	id := c.Param("id")
	logger.LogInfo("Invoked controller to fetch order with key" + id, c)
	c.JSON(http.StatusOK, Repo.FetchItemByKey(id))
}

func (ctrl *OrderController) FetchAllOrders(c *gin.Context) {
	logger.LogInfo("Invoked controller to fetch orders", c)
	c.JSON(http.StatusOK, Repo.FetchAllRecords())
}