package controller

import (
	"delivery_tracking_api/producer/logger"
	"delivery_tracking_api/producer/model"
	"delivery_tracking_api/producer/producer"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var orders = []model.Order{
	{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
}

func Init() {
	logger.Infoln("Ready to produce events")
	producer.Init()
	go producer.Log()
	producer.ProduceEvents(orders)
}

type OrderController struct{
	OrderControllerInterface
}

func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	logger.LogInfo("placed new order with id : "+order.OrderId, c)
	order.OrderDate = time.Now().Format("01-02-2006 15:04:05")
	order.Status.StatusDate = time.Now().Format("01-02-2006 15:04:05")
	orders = append(orders, order)
	c.JSON(http.StatusAccepted, "order successfully added with id : "+order.OrderId)
	producer.ProduceEvent(order)
}

func (ctrl *OrderController) GetPendingOrders(c *gin.Context) {
	var pendingOrders []model.Order
	for _, order := range orders {
		if strings.ToLower(order.Status.State) != "delivered" {
			pendingOrders = append(pendingOrders, order)
		}
	}
	logger.LogInfo("fetching pending orders", c)
	c.JSON(http.StatusOK, pendingOrders)
}

func (ctrl *OrderController) GetDeliveredOrders(c *gin.Context) {
	var deliveredOrders []model.Order
	for _, order := range orders {
		if strings.ToLower(order.Status.State) == "delivered" {
			deliveredOrders = append(deliveredOrders, order)
		}
	}
	logger.LogInfo("fetching delivered orders", c)
	c.JSON(http.StatusOK, deliveredOrders)
}

func (ctrl *OrderController) ChangeState(c *gin.Context) {
	state := c.Query("state")
	id := c.Query("id")
	for i, order := range orders {
		if order.OrderId == id {
			order.Status.State = state
			order.Status.StatusDate = time.Now().Format("01-02-2006 15:04:05")
			orders[i] = order
			producer.ProduceEvent(order)
		}
	}
	logger.LogInfo("state changed to "+state+", with id : "+id, c)
	c.JSON(http.StatusAccepted, "order state changed successfully with id : "+id)
}

func (ctrl *OrderController) ChangeAddress(c *gin.Context) {
	id := c.Param("id")
	var address model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for i, order := range orders {
		if order.OrderId == id {
			order.Address = address
			orders[i] = order
			producer.ProduceEvent(order)
		}
	}
	logger.LogInfo("Address changed to "+address.City+", "+address.State+", with id : "+id, c)
	c.JSON(http.StatusAccepted, "order address changed successfully with id : "+id)
}
