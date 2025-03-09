package router

import (
	"delivery_tracking_api/consumer2/consumer"
	"delivery_tracking_api/consumer2/controller"
	"delivery_tracking_api/consumer2/logger"

	"github.com/gin-gonic/gin"
)

func Init(){
	logger.Infoln("ready to consume events")
	consumer.Init()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	orderController := &controller.OrderController{}

	consumer := router.Group("/consumer") 
	{
		consumer.GET("/order/:id", orderController.GetOrder)
		consumer.GET("/status/:id", orderController.GetStatus)
		consumer.GET("/orders", orderController.GetAllOrders)
	}
	router.Run(":8082")
}