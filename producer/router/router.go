package router

import (
	"delivery_tracking_api/producer/controller"
	docs "delivery_tracking_api/producer/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	controller.Init()
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/producer"

	orderController := &controller.OrderController{}

	producer := router.Group("/producer")
	{
		producer.POST("/place", orderController.PlaceOrder)
		producer.GET("/pending", orderController.GetPendingOrders)
		producer.GET("/delivered", orderController.GetDeliveredOrders)
		producer.GET("/changeState", orderController.ChangeState)
		producer.POST("/changeAddress/:id", orderController.ChangeAddress)
		producer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	router.Run()
}
