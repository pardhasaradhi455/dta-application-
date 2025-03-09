package main

import (
	"delivery_tracking_api/consumer1/logger"
	"delivery_tracking_api/consumer1/router"
)

func main(){

	logger.Init()
	logger.Infoln("logger initialized")

	logger.Infoln("initializing router")
	router.Init()
}