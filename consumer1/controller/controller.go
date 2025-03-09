package controller

import (
	"delivery_tracking_api/consumer1/logger"
	"delivery_tracking_api/consumer1/model"
	"delivery_tracking_api/consumer1/repo"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddOrder(order model.Order) {
	Repo.InserOrUpdateRecord(order.OrderId, order)
}

type OrderController struct{
	OrderControllerInterface
}

var Repo = &repo.Repo{}

func (ctrl *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order := Repo.FetchItemByKey(id)
	j, _ := json.Marshal(order)
	logger.LogInfo("status fetched for order : "+string(j), c)
	c.JSON(http.StatusOK, order)
}


func (ctrl *OrderController) GetStatus(c *gin.Context) {
	id := c.Param("id")
	order := Repo.FetchItemByKey(id)
	var view string
	switch order.Status.State {
	case "placed":
		view = "placed.html"
	case "out-for-delivery":
		view = "out.html"
	case "delivered":
		view = "delivered.html"
	default:
		view = "shipped.html"
	}
	c.HTML(http.StatusOK,view, gin.H{
		"id" : id,
		"orderDate" : order.OrderDate,
		"status" : order.Status.State,
		"statusDate" : order.Status.StatusDate,
		"total" : order.OrderTotal,
		"city" : order.Address.City,
		"state" : order.Address.State,
		"pincode" : order.Address.Pincode,
	})
}

func (ctrl *OrderController) GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, Repo.FetchAllRecords())
}