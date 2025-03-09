package router

import (
	"delivery_tracking_api/consumer2DB/controller"
	"delivery_tracking_api/consumer2DB/logger"
	"delivery_tracking_api/consumer2DB/repo"

	"github.com/gin-gonic/gin"
)

func Init() {
	logger.Infoln("Ready to use Badger DB")
	repo.Init()
	router := gin.Default()

	OrderController := controller.OrderController{}

	db := router.Group("/db") 
	{
		db.POST("/insert", OrderController.InsertOrder)
		db.GET("/fetch/:id", OrderController.FetchOrderByKey)
		db.GET("/fetch", OrderController.FetchAllOrders)
	}
	router.Run(":7072")
}