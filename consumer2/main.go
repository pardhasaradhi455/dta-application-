package main

import (
	"delivery_tracking_api/consumer2/logger"
	"delivery_tracking_api/consumer2/router"
)

func main(){

	logger.Init()
	logger.Infoln("logger initialized")

	logger.Infoln("initializing router")
	router.Init()
}